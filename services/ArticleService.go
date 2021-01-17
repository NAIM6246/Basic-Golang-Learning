package services

import (
	"Golang/models"
	"Golang/repository"
)

//Atricle	:
type ArticleService struct {
	repository.ArticleRepository
}

func NewArticleRepository() *ArticleService {
	return &ArticleService{}
}

//Get Article by id
func (h *ArticleService) GetArticleByID(id uint) (*models.Article, error) {
	article := models.Article{}
	return &article, nil
}

//Create Article
func (h *ArticleService) CreateArticle(article *models.Article) (*models.Article, error) {
	return h.ArticleRepository.Create(article)
}
