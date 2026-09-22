package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

func (h *MeHandler) UpdateMe(w http.ResponseWriter, r *http.Request) {
	var request struct {
		DisplayName string `json:"display_name"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	displayName := strings.TrimSpace(request.DisplayName)

	if displayName == "" {
		http.Error(w, "Display name cannot be empty", http.StatusBadRequest)
		return
	}

	if len(displayName) > 100 {
		http.Error(w, "Display name is too long", http.StatusBadRequest)
		return
	}

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

	err = h.UserRepository.UpdateDisplayName(userID, displayName)
	if err != nil {
		http.Error(w, "Failed to update display name", http.StatusInternalServerError)
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

func (h *MeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetMe(w, r)

	case http.MethodPatch:
		h.UpdateMe(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
