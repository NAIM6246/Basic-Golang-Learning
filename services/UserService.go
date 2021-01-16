package services

import (
	"Golang/models"
	"Golang/repository"
)

//UserService structure
type UserService struct {
	repository.UserRepository
}

//NewUserService Constructor of UserService
func NewUserService() *UserService {
	return &UserService{}
}

//Geting user thorugh id
func (h *UserService) GetUserByID(id uint) (*models.User, error) {

	user := models.User{
		ID:   id,
		NAME: "naim",
	}
	return &user, nil
}

func (h *UserService) CreateUser(user *models.User) (*models.User, error) {
	return h.UserRepository.Create(user)
}
