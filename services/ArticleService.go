package services

import (
	"Golang/config"
	"Golang/conn"
	"Golang/models"
	"Golang/repository"
)

//Atricle	:
type ArticleService struct {
	articleRepository *repository.ArticleRepository
}

func NewArticleService() *ArticleService {
	con := conn.ConnectDB(config.NewDBConfig())
	return &ArticleService{
		articleRepository: repository.NewArticleRepository(con),
	}
}

//Get Article by id
func (h *ArticleService) GetArticleByID(id uint) (*models.Article, error) {
	article := models.Article{}
	return &article, nil
}

//Create Article
func (h *ArticleService) CreateArticle(article *models.Article) (*models.Article, error) {
	return h.articleRepository.Create(article)
}

//Get Article by id
func (h *ArticleService) GetArticle() ([]*models.Article, error) {
	return h.articleRepository.Get()
}
