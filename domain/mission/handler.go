package mission

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/blazeisclone/spaceops-mission-ctrl/domain"
)

type Handler struct {
	repo MissionRepository
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		repo: NewMySQLMissionRepository(db),
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mission := &domain.Mission{
		Name:        payload.Name,
		Description: payload.Description,
	}

	if err := h.repo.Create(mission); err != nil {
		http.Error(w, fmt.Sprintf("Error creating mission: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mission)
}

func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {
	mission, err := h.repo.Read()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding mission: %v", err), http.StatusInternalServerError)
		return
	}

	if mission == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mission)
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid mission ID", http.StatusBadRequest)
		return
	}

	mission, err := h.repo.FindByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding mission: %v", err), http.StatusInternalServerError)
		return
	}

	if mission == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mission)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid mission ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mission := &domain.Mission{
		ID:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
	}

	err = h.repo.UpdateByID(id, mission)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating mission: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Mission updated successfully")
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid mission ID", http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting mission: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Mission deleted successfully")
}
