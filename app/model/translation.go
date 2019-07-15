package model

// Represents a single Bible translation
type Translation struct {
	Id            int    `gorm:"primary_key" json:"id"`
	Table         string `gorm:"table" json:"-"`
	Abbreviation  string `gorm:"abbreviation" json:"abbreviation"`
	Language      string `gorm:"language" json:"language"`
	Version       string `gorm:"version" json:"translation"`
	InfoText      string `gorm:"info_text" json:"info_text"`
	InfoURL       string `gorm:"info_url" json:"info_url"`
	Publisher     string `gorm:"publisher" json:"publisher"`
	Copyright     string `gorm:"copyright" json:"copyright"`
	CopyrightInfo string `gorm:"copyright_info" json:"copyright_info"`
}

// Specify the DB table for GORM
func (Translation) TableName() string {
	return "bible_version_key"
}
