package project

import (
	"gorm.io/gorm"
	"github.com/seu-usuario/portfolio/applications/api/internal/models"
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