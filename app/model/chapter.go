package model

// Represents a single chapter
type Chapter struct {
	Chapter int `gorm:"Column:c" json:"chapter"`
}
