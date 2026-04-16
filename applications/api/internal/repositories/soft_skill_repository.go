package repositories

import (
	"gorm.io/gorm"
	"github.com/cirodourado/portfolio/applications/api/internal/models"
)

type SoftSkillRepository interface {
    GetAll() ([]models.SoftSkill, error)
}

type GormSoftSkillRepository struct {
	db *gorm.DB
}

func NewGormSoftSkillRepository(db *gorm.DB) *GormSoftSkillRepository {
    return &GormSoftSkillRepository{db: db}
}

func (gssr *GormSoftSkillRepository) GetAll() ([]models.SoftSkill, error) {
	var softSkills []models.SoftSkill

	result := gssr.db.Order("\"order\" asc").Find(&softSkills)

	return softSkills, result.Error
}