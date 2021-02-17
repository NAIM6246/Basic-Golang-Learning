package repository

import (
	"Golang/conn"
	"Golang/models"

	"github.com/jinzhu/gorm"
)

//
type UserRepository struct {
	db *gorm.DB
}

//
func NewUserRepository(db *conn.DB) *UserRepository {
	return &UserRepository{
		db: db.Table(models.UserTable()),
	}
}

//Create :
func (repo *UserRepository) Create(u *models.User) (*models.User, error) {
	/*fmt.Println("print 0")
	var err error
	if !repo.db.NewRecord(&u) {
		repo.db.Create(&u)
		if !repo.db.NewRecord(&u) {
			fmt.Println("print 1")
			return u, nil
		}
		fmt.Println("print 2")
		return nil, err
	}
	fmt.Println("print 3")
	return nil, err

	fmt.Println(u)
	*/
	if err := repo.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
