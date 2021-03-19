package services

import (
	"Golang/models"
	"Golang/repository"
)

//
type IUserService interface {
	GetAll() ([]*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

//UserService structure
type UserService struct {
	userRepository repository.IUserRepository
}

//NewUserService Constructor of UserService
func NewUserService(userRepository repository.IUserRepository) IUserService {
	//con := conn.ConnectDB(config.NewDBConfig())
	return &UserService{
		userRepository: userRepository,
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
