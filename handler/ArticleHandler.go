package handler

import (
	"Golang/services"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

//Article handler	:
type ArticleHandler struct {
	articleService services.ArticleService
}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{
		articleService: services.ArticleService{},
	}
}

func (h *ArticleHandler) Handler(rout chi.Router) {
	rout.Get("/get", h.getArticle)
	rout.Post("/post", h.createArticle)
}

func (h *ArticleHandler) getArticle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("message posted")
}

func (h *ArticleHandler) createArticle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("article created")
}
