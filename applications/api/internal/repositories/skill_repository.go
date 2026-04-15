package repositories

import (
	"github.com/cirodourado/portfolio/applications/api/internal/models"
	"gorm.io/gorm"
)

type SkillRepository interface {
	GetAllWithSkills() ([]models.Category, error)
}

type GormSkillRepository struct {
	db *gorm.DB
}

func NewGormSkillRepository(db *gorm.DB) *GormSkillRepository {
	return &GormSkillRepository{db: db}
}

func (gsr *GormSkillRepository) GetAllWithSkills() ([]models.Category, error) {
	var categories []models.Category

	result := gsr.db.
		Preload("Skills", func(db *gorm.DB) *gorm.DB {
			return db.Order("proficiency_level DESC, name ASC")
		}).
		Order("\"order\" ASC, name ASC").
		Find(&categories)

	return categories, result.Error
}