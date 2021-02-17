package services

import (
	"Golang/config"
	"Golang/conn"
	"Golang/models"
	"Golang/repository"
)

//UserService structure
type UserService struct {
	userRepository *repository.UserRepository
}

//NewUserService Constructor of UserService
func NewUserService() *UserService {
	con := conn.ConnectDB(config.NewDBConfig())
	return &UserService{
		userRepository: repository.NewUserRepository(con),
	}
}

//Geting user thorugh id
func (h *UserService) GetUserByID(id uint) (*models.User, error) {

	var user models.User
	return &user, nil
}

func (h *UserService) CreateUser(user *models.User) (*models.User, error) {
	return h.userRepository.Create(user)
}

//Get all user :
func (h *UserService) GetAll() ([]*models.User, error) {
	return h.userRepository.GetAll()
}
