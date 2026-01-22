package project

import "github.com/cirodourado/portfolio/applications/api/internal/models"

type ProjectRepository interface {
	GetAll() ([]models.Project, error)
}