package models

//User model
type User struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	NAME     string    `gorm:"typevarchar(100);not null" json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Articles []Article `gorm:"foreign_key:UserID"`
}

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:password`
}

//TableName for User
func UserTable() string {
	return "Users"
}
