package config

import "sync"

//DB :
type DBConfig struct {
	Server   string
	Port     int
	User     string
	Password string
	DbName   string
}

var dbConfig *DBConfig

func mapDbConfig() {
	dbConfig = &DBConfig{
		Server:   "localhost",
		Port:     3000,
		User:     "",
		Password: "",
		DbName:   "BasicGoLang",
	}
}

//returning dbconfig instance :
func NewDBConfig() *DBConfig {
	var loadDBOnce sync.Once
	loadDBOnce.Do(mapDbConfig)
	return dbConfig
}
