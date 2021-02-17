package models

//Article model
type Article struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Title    string `gorm:"type=varchar(300);not null" json:"title"`
	Body     string `gorm:"type=text" json:"body"`
	Author   User   `gorm:"foreignkey:AuthorID" json:"author"`
	AuthorID uint   `json:"authorId"`
}

//TableName for Article
func ArticleTable() string {
	return "articles"
}
