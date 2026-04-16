package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cirodourado/portfolio/applications/api/internal/repositories"
)

type SoftSkillHandler struct {
	softSkillRepository repositories.SoftSkillRepository
}

func NewSoftSkillHandler(softSkillRepository repositories.SoftSkillRepository) *SoftSkillHandler {
	return &SoftSkillHandler{softSkillRepository: softSkillRepository}
}

func (ssh *SoftSkillHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	experiences, err := ssh.softSkillRepository.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(experiences)
}