package project

import (
	"gorm.io/gorm"
	"github.com/cirodourado/portfolio/applications/api/internal/models"
)

type GormProjectRepository struct {
	DB *gorm.DB
}

func NewGormProjectRepository(db *gorm.DB) *GormProjectRepository {
	return &GormProjectRepository{DB: db}
}

func (r *GormProjectRepository) GetAll() ([]models.Project, error) {
	var projects []models.Project

	result := r.DB.Preload("Skills.Category").Find(&projects)

	return projects, result.Error
}

func (r *GormProjectRepository) GetBySlug(slug string) (*models.Project, error) {
	var project models.Project

	result := r.DB.Preload("Skills.Category").Where("slug = ?", slug).First(&project)

	if result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}