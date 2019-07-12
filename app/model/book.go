package model

type Book struct {
	Id   int    `gorm:"Column:b" json:"id"`
	Name string `gorm:"Column:n" json:"name"`
}
