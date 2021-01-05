package services

//struct :
type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

//Geting user id :
func (h *UserService) GetUserByID(uint) (string, error) {
	return "Naim", nil
}
