package model

type Chapter struct {
	Chapter int `gorm:"Column:c" json:"chapter"`
}
