package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cirodourado/portfolio/applications/api/internal/repositories"
)

type SkillHandler struct {
	repo repositories.SkillRepository
}

func NewSkillHandler(repo repositories.SkillRepository) *SkillHandler {
	return &SkillHandler{repo: repo}
}

func (sh *SkillHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := sh.repo.GetAllWithSkills()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}