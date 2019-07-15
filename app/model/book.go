// The model package provides data models for various parts of The Bible.
package model

// Represents a single book
type Book struct {
	Id   int    `gorm:"Column:b" json:"id"`
	Name string `gorm:"Column:n" json:"name"`
}
