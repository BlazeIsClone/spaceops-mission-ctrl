package organization

import (
	"net/http"
	"os"
)

func Routes(router *http.ServeMux) {
	handler := &Handler{}

	basePath := os.Getenv("BASE_PATH")

	router.HandleFunc("POST "+basePath+"/organizations", handler.Create)
	router.HandleFunc("GET "+basePath+"/organizations/{id}", handler.FindByID)
	router.HandleFunc("PUT "+basePath+"/organizations/{id}", handler.UpdateByID)
	router.HandleFunc("DELETE "+basePath+"/organizations/{id}", handler.DeleteByID)
}
