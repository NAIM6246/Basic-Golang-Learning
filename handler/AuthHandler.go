package handler

import (
	"Golang/auth"
	"Golang/models"
	"Golang/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

//
type IAuthHandler interface {
	IHandler
}

type AuthHandler struct {
	authService services.IAuthService
	auth        auth.IAuth
}

func NewAuthhandler(authService services.IAuthService,
	auth auth.IAuth) IAuthHandler {
	return &AuthHandler{
		authService: authService,
		auth:        auth,
	}
}

func (h *AuthHandler) Handle(router chi.Router) {
	//fmt.Println("auth")
	router.Post("/login", h.login)
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	var loginObj models.UserLoginDto
	parsingErr := json.NewDecoder(r.Body).Decode(&loginObj)
	if parsingErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(parsingErr)
		return
	}
	token, err := h.authService.Login(&loginObj)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println("Logged in")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(token)
}
