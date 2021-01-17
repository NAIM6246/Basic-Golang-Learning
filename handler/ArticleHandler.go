package handler

import (
	"Golang/models"
	"Golang/services"
	"encoding/json"
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
	article := models.Article{}
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		panic(err)
	}
	d, e := h.articleService.Create(&article)
	if e != nil {
		panic(e)
		return
	}
	fmt.Println(w, d)
}
