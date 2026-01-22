package project

import "github.com/seu-usuario/portfolio/applications/api/internal/models"

type ProjectRepository interface {
	GetAll() ([]models.Project, error)
}