package handler

import (
	"bible/app/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetAllVersions(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	versions := []model.Version{}
	db.Find(&versions)
	respondJSON(w, http.StatusOK, versions)
}
