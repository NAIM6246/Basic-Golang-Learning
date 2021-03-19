package services

import (
	"Golang/models"
	"Golang/repository"
)

//
type IArticleService interface {
	CreateArticle(article *models.Article) (*models.Article, error)
	GetArticle() ([]*models.Article, error)
}

//Atricle	:
type ArticleService struct {
	articleRepository repository.IArticleRepository
}

func NewArticleService(articleRepository repository.IArticleRepository) IArticleService {
	//con := conn.ConnectDB(config.NewDBConfig())
	return &ArticleService{
		articleRepository: articleRepository,
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
