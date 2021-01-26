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
	articleService *services.ArticleService
}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{
		articleService: services.NewArticleService(),
	}
}

func (h *ArticleHandler) Handler(rout chi.Router) {
	rout.Get("/", h.getArticle)
	rout.Post("/", h.createArticle)
}

func (h *ArticleHandler) getArticle(w http.ResponseWriter, req *http.Request) {
	//this function is not completed yet
	fmt.Fprintf(w, "message posted")
}

func (h *ArticleHandler) createArticle(w http.ResponseWriter, req *http.Request) {
	article := models.Article{}
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		//bad request error
		w.WriteHeader(400)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message" : "bad request error"}`))
		return
	}
	fmt.Println(article)
	d, e := h.articleService.CreateArticle(&article)
	if e != nil {
		//bad request error
		w.WriteHeader(400)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message" : "bad request error"}`))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	_ = json.NewEncoder(w).Encode(d)
}
