package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cirodourado/portfolio/applications/api/internal/repositories"
)

type ProfileHandler struct {
    profileRepository repositories.ProfileRepository
}

func NewProfileHandler(profileRepository repositories.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{profileRepository: profileRepository}
}

func (ph *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	profile, err := ph.profileRepository.GetProfile()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}