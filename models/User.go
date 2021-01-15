package models

//User model
type User struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	NAME     string    `gorm:"typevarchar(100);not null" json:"name"`
	Articles []Article `gorm:"foreign_key:UserID"`
}

//TableName for User
func (User) TableName() string {
	return "Users"
}
