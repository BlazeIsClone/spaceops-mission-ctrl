package mission

import (
	"database/sql"
	"net/http"
	"os"
)

func Routes(router *http.ServeMux, db *sql.DB) {
	handler := NewHandler(db)

	basePath := os.Getenv("BASE_PATH")

	router.HandleFunc("GET "+basePath+"/missions", handler.Index)
	router.HandleFunc("POST "+basePath+"/missions", handler.Store)
	router.HandleFunc("GET "+basePath+"/missions/{id}", handler.Show)
	router.HandleFunc("PUT "+basePath+"/missions/{id}", handler.Update)
	router.HandleFunc("DELETE "+basePath+"/missions/{id}", handler.Destroy)
}
