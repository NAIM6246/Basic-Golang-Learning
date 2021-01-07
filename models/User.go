package models

//User :
type User struct {
	ID   uint
	NAME string
}

//Users	:
func (User) TableName() string {
	return "User"
}
