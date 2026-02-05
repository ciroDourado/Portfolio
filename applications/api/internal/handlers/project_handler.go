package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

    "github.com/go-chi/chi/v5"
	"github.com/cirodourado/portfolio/applications/api/internal/repositories"
)

type ProjectHandler struct {
	projectRepository repositories.ProjectRepository
    storageURL        *url.URL
}

func NewProjectHandler(projectRepository repositories.ProjectRepository, storageURL *url.URL) *ProjectHandler {
	return &ProjectHandler{projectRepository: projectRepository, storageURL: storageURL}
}

func (ph *ProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	projects, err := ph.projectRepository.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    for i := range projects {
        if projects[i].ThumbnailURL != nil {
            thumbnailURL := ph.storageURL.JoinPath(*projects[i].ThumbnailURL).String()
            projects[i].ThumbnailURL = &thumbnailURL
        }
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (ph *ProjectHandler) GetBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		http.Error(w, "Slug is required", http.StatusBadRequest)
		return
	}

	project, err := ph.projectRepository.GetBySlug(slug)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

    if project.ThumbnailURL != nil {
	    thumbnailURL := ph.storageURL.JoinPath(*project.ThumbnailURL).String()
	    project.ThumbnailURL = &thumbnailURL
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}