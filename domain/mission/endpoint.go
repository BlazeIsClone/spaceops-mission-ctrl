package mission

import (
	"database/sql"
	"net/http"
	"os"
)

func Routes(router *http.ServeMux, db *sql.DB) {
	handler := NewHandler(db)

	basePath := os.Getenv("BASE_PATH")

	router.HandleFunc("POST "+basePath+"/missions", handler.Create)
	router.HandleFunc("GET "+basePath+"/missions/{id}", handler.FindByID)
	router.HandleFunc("PUT "+basePath+"/missions/{id}", handler.UpdateByID)
	router.HandleFunc("DELETE "+basePath+"/missions/{id}", handler.DeleteByID)
}
