package repository

import (
	"Golang/conn"
	"Golang/models"

	"github.com/jinzhu/gorm"
)

//Article	:
type ArticleRepository struct {
	db *gorm.DB
}

//Articlerepository	:
func NewArticleRepository(db *conn.DB) *UserRepository {
	return &UserRepository{
		db: db.Table(models.ArticleTable()),
	}
}

//Create article	:
func (repo *ArticleRepository) Create(art *models.Article) (*models.Article, error) {
	var err error
	if !repo.db.NewRecord(art) {
		repo.db.Create(&art)
		if !repo.db.NewRecord(&art) {
			return art, nil
		}
		return nil, err
	}
	return nil, err
}
