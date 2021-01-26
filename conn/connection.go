package conn

import (
	"Golang/config"
	"Golang/models"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	//mssql
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

const (
	dbtype = "mssql"
)

//DB structure :
type DB struct {
	*gorm.DB
}

func NewDB() *DB {
	return &DB{}
}

var dbInstance *DB

var connDBOnce sync.Once

func ConnectDB(config *config.DBConfig) *DB {
	//var connDBOnce sync.Once
	connDBOnce.Do(func() {
		_ = connectDB(config)
	})

	//dbInstance.Migration()
	return dbInstance
}

func connectDB(config *config.DBConfig) error {
	connString := fmt.Sprintf("server=%s; port=%d; database=%s;", config.Server, config.Port, config.DbName)
	conn, err := gorm.Open(dbtype, connString)
	if err != nil {
		log.Fatal("Open connection faied: ", err.Error())
		return err
	}
	fmt.Println("Databse is connected succesfully")
	dbInstance = &DB{conn}
	return nil
}

func (db *DB) Migration() {
	db.AutoMigrate(&models.User{}, &models.Article{})
}
