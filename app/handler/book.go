package handler

import (
	"bible/app/model"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var books_map = make(map[string][]model.Book)

func GetAllBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	translation_id, _ := vars["translation_id"]

	if nil == books_map[translation_id] {
		var translation model.Translation
		db.Raw("SELECT * FROM bible_version_key WHERE id = ?", translation_id).Scan(&translation)
		table_name := translation.Table
		rows, _ := db.Raw("SELECT * FROM key_english WHERE b IN (SELECT DISTINCT(b) FROM " + table_name + ") ORDER BY b ASC").Rows()
		defer rows.Close()
		for rows.Next() {
			var book model.Book
			db.ScanRows(rows, &book)
			books_map[translation_id] = append(books_map[translation_id], book)
		}
	}

	respondJSON(w, http.StatusOK, books_map[translation_id])
}
