package repository

import (
	"Golang/conn"
	"Golang/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Article	:
type ArticleRepository struct {
	db *gorm.DB
}

//Articlerepository	:
func NewArticleRepository(db *conn.DB) *ArticleRepository {
	return &ArticleRepository{
		db: db.Table(models.ArticleTable()),
	}
}

//Create article	:
func (repo *ArticleRepository) Create(art *models.Article) (*models.Article, error) {
	if err := repo.db.Create(&art).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return art, nil
}

func (repo *ArticleRepository) Get() ([]*models.Article, error) {
	var art []*models.Article
	err := repo.db.Find(&art).Error
	if err != nil {
		return nil, err
	}
	return art, nil
}
