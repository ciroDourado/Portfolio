package handlers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
	"github.com/cirodourado/portfolio/applications/api/internal/models"
)

type ExperienceHandler struct {
	DB *gorm.DB
}

func NewExperienceHandler(db *gorm.DB) *ExperienceHandler {
	return &ExperienceHandler{DB: db}
}

// GetAll (GET /api/v1/experiences)
func (h *ExperienceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var experiences []models.Experience

	result := h.DB.Order("start_date desc").Find(&experiences)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(experiences)
}