package experience

import "github.com/cirodourado/portfolio/applications/api/internal/models"

type ExperienceRepository interface {
	GetAll() ([]models.Experience, error)
}