package model

type Version struct {
	Id            int    `gorm:"primary_key" json:"id"`
	Table         string `gorm:"table" json:"table"`
	Abbreviation  string `gorm:"abbreviation" json:"abbreviation"`
	Language      string `gorm:"language" json:"language"`
	Version       string `gorm:"version" json:"version"`
	InfoText      string `gorm:"info_text" json:"info_text"`
	InfoURL       string `gorm:"info_url" json:"info_url"`
	Publisher     string `gorm:"publisher" json:"publisher"`
	Copyright     string `gorm:"copyright" json:"copyright"`
	CopyrightInfo string `gorm:"copyright_info" json:"copyright_info"`
}

func (Version) TableName() string {
	return "bible_version_key"
}
