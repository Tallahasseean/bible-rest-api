package handler

import (
	"bible/app/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Slice of translations to store in memory, reducing multiple lookups.
var translations []model.Translation

// Returns all translations available in this API.
func GetAllTranslations(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if len(translations) == 0 {
		translations = []model.Translation{}
		db.Find(&translations)
	}
	respondJSON(w, http.StatusOK, translations)
}
