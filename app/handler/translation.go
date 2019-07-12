package handler

import (
	"bible/app/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetAllTranslations(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	translations := []model.Translation{}
	db.Find(&translations)
	respondJSON(w, http.StatusOK, translations)
}
