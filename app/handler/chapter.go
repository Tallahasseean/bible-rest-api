package handler

import (
	"bible/app/model"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var chapters_map = make(map[string][]model.Chapter)

func GetAllChapters(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	translation_id, _ := vars["translation_id"]
	book_id, _ := vars["book_id"]

	if nil == chapters_map[translation_id+book_id] {
		var translation model.Translation
		db.Raw("SELECT * FROM bible_version_key WHERE id = ?", translation_id).Scan(&translation)
		table_name := translation.Table
		rows, _ := db.Raw("SELECT c FROM "+table_name+" WHERE b = ? GROUP BY c", book_id).Rows()
		defer rows.Close()
		for rows.Next() {
			var chapter model.Chapter
			db.ScanRows(rows, &chapter)
			chapters_map[translation_id+book_id] = append(chapters_map[translation_id+book_id], chapter)
		}
	}

	respondJSON(w, http.StatusOK, chapters_map[translation_id+book_id])
}
