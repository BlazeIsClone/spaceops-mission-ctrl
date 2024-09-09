package organization

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("created\n"))
}

func (h *Handler) FindByID(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("found: " + req.PathValue("id") + "\n"))
}

func (h *Handler) UpdateByID(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("updated: " + req.PathValue("id") + "\n"))
}

func (h *Handler) DeleteByID(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("deleted: " + req.PathValue("id") + "\n"))
}
