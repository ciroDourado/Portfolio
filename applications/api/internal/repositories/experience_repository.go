package repositories

import (
	"gorm.io/gorm"
	"github.com/cirodourado/portfolio/applications/api/internal/models"
)

type ExperienceRepository interface {
	GetAll() ([]models.Experience, error)
}

type GormExperienceRepository struct {
	db *gorm.DB
}

func NewGormExperienceRepository(db *gorm.DB) *GormExperienceRepository {
	return &GormExperienceRepository{db: db}
}

func (ger *GormExperienceRepository) GetAll() ([]models.Experience, error) {
	var experiences []models.Experience

	result := ger.db.Order("start_date desc").Find(&experiences)

	return experiences, result.Error
}
