package handler

import (
	"bible/app/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

var translations []model.Translation

func GetAllTranslations(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if len(translations) == 0 {
		translations = []model.Translation{}
		db.Find(&translations)
	}
	respondJSON(w, http.StatusOK, translations)
}
