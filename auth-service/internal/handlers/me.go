package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/middleware"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/repository"
)

type MeHandler struct {
	UserRepository *repository.UserRepository
}

func NewMeHandler(userRepository *repository.UserRepository) *MeHandler {
	return &MeHandler{
		UserRepository: userRepository,
	}
}

func (h *MeHandler) GetMe(w http.ResponseWriter, r *http.Request) {

	userIDValue := r.Context().Value(middleware.UserIDKey)

	userIDString, ok := userIDValue.(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	user, err := h.UserRepository.FindUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
