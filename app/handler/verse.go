package handler

import (
	"bible/app/model"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var verses_map = make(map[string][]model.Verse)

func GetAllVerses(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	translation_id, _ := vars["translation_id"]
	book_id, _ := vars["book_id"]
	chapter, _ := vars["chapter"]

	if nil == verses_map[translation_id+book_id+chapter] {
		var translation model.Translation
		db.Raw("SELECT * FROM bible_version_key WHERE id = ?", translation_id).Scan(&translation)
		table_name := translation.Table
		rows, _ := db.Raw("SELECT v FROM "+table_name+" WHERE b = ? AND c = ?", book_id, chapter).Rows()
		defer rows.Close()
		for rows.Next() {
			var verse model.Verse
			db.ScanRows(rows, &verse)
			verses_map[translation_id+book_id+chapter] = append(verses_map[translation_id+book_id+chapter], verse)
		}
	}

	respondJSON(w, http.StatusOK, verses_map[translation_id+book_id+chapter])
}
