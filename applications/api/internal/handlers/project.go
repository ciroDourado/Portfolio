package handlers

import (
	"encoding/json"
	"net/http"

	repositories "github.com/seu-usuario/portfolio/applications/api/internal/repositories/project"
)

type ProjectHandler struct {
	Repo repositories.ProjectRepository
}

func NewProjectHandler(repo repositories.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{Repo: repo}
}

func (h *ProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Repo.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}