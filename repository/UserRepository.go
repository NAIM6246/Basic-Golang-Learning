package repository

import (
	"Golang/conn"
	"Golang/models"

	"github.com/jinzhu/gorm"
)

//
type UserRepository struct {
	db *gorm.DB
}

//
func NewUserRepository(db *conn.DB) *UserRepository {
	return &UserRepository{
		db: db.Table(models.UserTable()),
	}
}

//Create :
func (repo *UserRepository) Create(u *models.User) (*models.User, error) {
	var err error
	if !repo.db.NewRecord(u) {
		repo.db.Create(&u)
		if !repo.db.NewRecord(u) {
			return u, nil
		}
		return nil, err
	}
	return nil, err
}
