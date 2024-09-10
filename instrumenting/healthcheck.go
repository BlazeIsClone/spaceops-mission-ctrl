package instrumenting

import (
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func Routes(router *http.ServeMux) {
	router.HandleFunc("GET /health", healthCheckHandler)
}
