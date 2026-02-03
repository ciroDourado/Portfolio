package experience

import (
	"gorm.io/gorm"
	"github.com/cirodourado/portfolio/applications/api/internal/models"
)

type GormExperienceRepository struct {
	DB *gorm.DB
}

func NewGormExperienceRepository(db *gorm.DB) *GormExperienceRepository {
	return &GormExperienceRepository{DB: db}
}

func (r *GormExperienceRepository) GetAll() ([]models.Experience, error) {
	var experiences []models.Experience

	result := r.DB.Order("start_date desc").Find(&experiences)

	return experiences, result.Error
}