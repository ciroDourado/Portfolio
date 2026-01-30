package handlers

import (
	"encoding/json"
	"net/http"

    "github.com/go-chi/chi/v5"
	repositories "github.com/cirodourado/portfolio/applications/api/internal/repositories/project"
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

func (h *ProjectHandler) GetBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		http.Error(w, "Slug is required", http.StatusBadRequest)
		return
	}

	project, err := h.Repo.GetBySlug(slug)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}