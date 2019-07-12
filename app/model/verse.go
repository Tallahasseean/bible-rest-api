package model

type Verse struct {
	Verse int `gorm:"Column:v" json:"verse"`
}
