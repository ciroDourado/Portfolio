package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cirodourado/portfolio/applications/api/internal/repositories"
)

type ProfileHandler struct {
    profileRepository repositories.ProfileRepository
    storageURL        *url.URL
}

func NewProfileHandler(profileRepository repositories.ProfileRepository, storageURL *url.URL) *ProfileHandler {
	return &ProfileHandler{profileRepository: profileRepository, storageURL: storageURL}
}

func (ph *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	profile, err := ph.profileRepository.GetProfile()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    if profile.AvatarURL != nil {
	    avatarURL := ph.storageURL.JoinPath(*profile.AvatarURL).String()
	    profile.AvatarURL = &avatarURL
    }
    if profile.CVURL != nil {
	    CVURL := ph.storageURL.JoinPath(*profile.CVURL).String()
	    profile.CVURL = &CVURL
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}