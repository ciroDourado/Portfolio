package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cirodourado/portfolio/applications/api/internal/repositories"
)

type ExperienceHandler struct {
	experienceRepository repositories.ExperienceRepository
}

func NewExperienceHandler(experienceRepository repositories.ExperienceRepository) *ExperienceHandler {
	return &ExperienceHandler{experienceRepository: experienceRepository}
}

func (eh *ExperienceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	experiences, err := eh.experienceRepository.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(experiences)
}