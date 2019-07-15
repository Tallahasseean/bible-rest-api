package model

// Represents a single verse
type Verse struct {
	Verse int    `gorm:"Column:v" json:"verse"`
	Text  string `gorm:"Column:t" json:"text,omitempty"`
}
