package services

import (
	"Golang/auth"
	"Golang/models"
	"Golang/repository"
	"errors"
	"fmt"
)

type IAuthService interface {
	Login(login *models.UserLoginDto) (string, error)
}

type AuthService struct {
	userRepository repository.IUserRepository
	auth           auth.IAuth
}

func NewAuthService(userRepository repository.IUserRepository,
	auth auth.IAuth) IAuthService {
	return &AuthService{
		userRepository: userRepository,
		auth:           auth,
	}
}

func (h *AuthService) Login(login *models.UserLoginDto) (string, error) {
	user, err := h.userRepository.Get(login)
	if err != nil {
		fmt.Println(err)
		fmt.Println("auth service")
		return "", err
	}
	if login.Password != user.Password {
		return "", errors.New("incorrect password")
	}
	token, err := h.auth.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil

}
