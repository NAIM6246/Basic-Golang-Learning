package repository

import "github.com/jinzhu/gorm"

type BaseRepository struct {
	db *gorm.DB
}
