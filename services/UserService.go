package services

import "Golang/models"

//struct :
type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

//Geting user id :
func (h *UserService) GetUserByID(a uint) (*models.User, error) {

	user := models.User{
		ID:   a,
		NAME: "naim",
	}
	return &user, nil
}
