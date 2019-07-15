package model

type Verse struct {
	Verse int    `gorm:"Column:v" json:"verse"`
	Text  string `gorm:"Column:t" json:"text,omitempty"`
}
