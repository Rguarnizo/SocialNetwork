package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rguarnizo/SocialMedia/pkg/auth"
	"github.com/rguarnizo/SocialMedia/profile-service/internal/service"
)

type ProfileHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *ProfileHandler {
	return &ProfileHandler{s}
}

func (h *ProfileHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID := auth.UserFromRequest(r)

	user, err := h.service.GetProfile(userID.ID)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
