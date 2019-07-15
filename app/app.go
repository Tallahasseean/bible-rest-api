package app

import (
	"bible/app/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

// Custom data type represents the core components of the application
type App struct {
	Environment string
	Router      *mux.Router
	DB          *gorm.DB
}

// Initialize the core parts of the application (env, DB, and router).
func (a *App) Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	a.Environment = os.Getenv("ENVIRONMENT")

	db, err := gorm.Open("sqlite3", "bible-sqlite.db")
	if err != nil {
		panic("Failed to connect database")
	}

	a.DB = db
	a.Router = mux.NewRouter()
	a.setRouters()
	a.Router.NotFoundHandler = http.HandlerFunc(NotFound)
}

// Decalre all routes and their corresponding handlers.
func (a *App) setRouters() {
	a.Get("/translations", a.handleRequest(handler.GetAllTranslations))
	a.Get("/translations/{translation_id:[0-9]}/books", a.handleRequest(handler.GetAllBooks))
	a.Get("/translations/{translation_id:[0-9]}/books/{book_id:[0-9]+}/chapters", a.handleRequest(handler.GetAllChapters))
	a.Get("/translations/{translation_id:[0-9]}/books/{book_id:[0-9]+}/chapters/{chapter:[0-9]+}/verses", a.handleRequest(handler.GetAllVerses))
	a.Get("/translations/{translation_id:[0-9]}/books/{book_id:[0-9]+}/chapters/{chapter:[0-9]+}/verses/{verse:[0-9]+-?[0-9]*}", a.handleRequest(handler.GetVerse))
}

// Convenience wrapper for GET requests.
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Run the server.
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// Custom type for a basic request handler function.
type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

// Convenience wrapper for request handlers.
func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	response := ("{\"error\":\"Resource not found\"}")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(response))
}
