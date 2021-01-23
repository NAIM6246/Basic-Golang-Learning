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
	rout.Get("/", h.getArticle)
	rout.Post("/", h.createArticle)
}

func (h *ArticleHandler) getArticle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "message posted")
}

func (h *ArticleHandler) createArticle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "article created")
	article := models.Article{}
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte("Name : "))
		panic(err)
	}
	d, e := h.articleService.Create(&article)
	if e != nil {
		panic(e)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("Name : " + d.Author.NAME))
	fmt.Fprintf(w, d.Author.NAME)
}
