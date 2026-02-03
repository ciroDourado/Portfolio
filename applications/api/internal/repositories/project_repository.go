package repositories

import (
	"gorm.io/gorm"
	"github.com/cirodourado/portfolio/applications/api/internal/models"
)

type ProjectRepository interface {
	GetAll() ([]models.Project, error)
	GetBySlug(slug string) (*models.Project, error)
}

type GormProjectRepository struct {
	db *gorm.DB
}

func NewGormProjectRepository(db *gorm.DB) *GormProjectRepository {
	return &GormProjectRepository{db: db}
}

func (gpr *GormProjectRepository) GetAll() ([]models.Project, error) {
	var projects []models.Project

	result := gpr.db.Preload("Skills.Category").Find(&projects)

	return projects, result.Error
}

func (gpr *GormProjectRepository) GetBySlug(slug string) (*models.Project, error) {
	var project models.Project

	result := gpr.db.Preload("Skills.Category").Where("slug = ?", slug).First(&project)

	if result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}
