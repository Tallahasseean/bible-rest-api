package handler

import (
	"bible/app/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var verses_map = make(map[string][]model.Verse)

func GetAllVerses(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var translation model.Translation
	vars := mux.Vars(r)
	translation_id, _ := vars["translation_id"]
	book_id, _ := vars["book_id"]
	chapter, _ := vars["chapter"]
	table_name := ""

	if nil == verses_map[translation_id+book_id+chapter] {
		if nil == translations_map[translation_id] {
			db.Raw("SELECT * FROM bible_version_key WHERE id = ?", translation_id).Scan(&translation)
			table_name = translation.Table
			translations_map[translation_id] = append(translations_map[translation_id], translation)
		} else {
			table_name = translations_map[translation_id][0].Table
		}
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

func GetVerse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var translation model.Translation
	vars := mux.Vars(r)
	translation_id, _ := vars["translation_id"]
	book_id, _ := vars["book_id"]
	chapter, _ := vars["chapter"]
	verse, _ := vars["verse"]
	verse_upper := ""
	table_name := ""

	if strings.Contains(verse, "-") {
		verse_upper = verse[strings.Index(verse, "-")+1:]
		verse = verse[:strings.Index(verse, "-")]

		verse_int, _ := strconv.Atoi(verse)
		verse_upper_int, _ := strconv.Atoi(verse_upper)

		if verse_int > verse_upper_int {
			respondError(w, http.StatusBadRequest, "The range of verses you requested appear to be transposed(?)")
			return
		}

		if verse_upper_int-verse_int >= 10 {
			respondError(w, http.StatusBadRequest, "Please limit your query to 10 verses per request")
			return
		}
	}

	if nil == translations_map[translation_id] {
		db.Raw("SELECT * FROM bible_version_key WHERE id = ?", translation_id).Scan(&translation)
		table_name = translation.Table
		translations_map[translation_id] = append(translations_map[translation_id], translation)
	} else {
		table_name = translations_map[translation_id][0].Table
	}
	rows, _ := db.Raw("SELECT v,t FROM "+table_name+" WHERE b = ? AND c = ? AND v >= ? AND v <= ?", book_id, chapter, verse, verse_upper).Rows()
	defer rows.Close()
	var result []model.Verse
	for rows.Next() {
		var verse model.Verse
		db.ScanRows(rows, &verse)
		result = append(result, verse)
	}

	respondJSON(w, http.StatusOK, result)
}
