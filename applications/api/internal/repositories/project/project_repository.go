package project

import "github.com/cirodourado/portfolio/applications/api/internal/models"

type ProjectRepository interface {
	GetAll() ([]models.Project, error)
	GetBySlug(slug string) (*models.Project, error)
}