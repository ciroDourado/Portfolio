package repositories

import (
	"gorm.io/gorm"
	"github.com/cirodourado/portfolio/applications/api/internal/models"
)

type ProfileRepository interface {
	GetProfile() (*models.Profile, error)
}

type GormProfileRepository struct {
	db *gorm.DB
}

func NewGormProfileRepository(db *gorm.DB) *GormProfileRepository {
	return &GormProfileRepository{db: db}
}

func (gpr *GormProfileRepository) GetProfile() (*models.Profile, error) {
	var profile models.Profile

	result := gpr.db.First(&profile)

	if result.Error != nil {
		return nil, result.Error
	}

	return &profile, nil
}
