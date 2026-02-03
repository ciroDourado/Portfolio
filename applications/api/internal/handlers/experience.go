package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cirodourado/portfolio/applications/api/internal/models"
	repositories "github.com/cirodourado/portfolio/applications/api/internal/repositories/experience"
)

type ExperienceHandler struct {
	Repo repositories.ExperienceRepository
}

func NewExperienceHandler(repo repositories.ExperienceRepository) *ExperienceHandler {
	return &ExperienceHandler{Repo: repo}
}

func (h *ExperienceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	experiences, err := h.Repo.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(experiences)
}