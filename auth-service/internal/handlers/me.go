package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/middleware"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/repository"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/storage"
)

type MeHandler struct {
	UserRepository *repository.UserRepository
	Storage        *storage.S3Storage
}

func NewMeHandler(
	userRepository *repository.UserRepository,
	storageService *storage.S3Storage,
) *MeHandler {
	return &MeHandler{
		UserRepository: userRepository,
		Storage:        storageService,
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

func (h *MeHandler) GetProfileImage(w http.ResponseWriter, r *http.Request) {
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

	if user.ProfileImageKey == "" {
		http.Error(w, "Profile image not found", http.StatusNotFound)
		return
	}

	object, err := h.Storage.Download(
		r.Context(),
		user.ProfileImageKey,
	)
	if err != nil {
		http.Error(w, "Failed to retrieve profile image", http.StatusInternalServerError)
		return
	}
	defer object.Body.Close()

	if object.ContentType != nil {
		w.Header().Set("Content-Type", *object.ContentType)
	}

	_, err = io.Copy(w, object.Body)
	if err != nil {
		http.Error(w, "Failed to send profile image", http.StatusInternalServerError)
		return
	}
}

func (h *MeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/me":
		h.GetMe(w, r)

	case r.Method == http.MethodPatch && r.URL.Path == "/me":
		h.UpdateMe(w, r)

	case r.Method == http.MethodPost && r.URL.Path == "/me/profile-image":
		h.UploadProfileImage(w, r)

	case r.Method == http.MethodGet && r.URL.Path == "/me/profile-image":
		h.GetProfileImage(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *MeHandler) UploadProfileImage(w http.ResponseWriter, r *http.Request) {
	// Limit the request body to 5 MB.
	r.Body = http.MaxBytesReader(w, r.Body, 5<<20)

	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		http.Error(w, "Image is too large or invalid", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("profile_image")
	if err != nil {
		http.Error(w, "Profile image is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check that the uploaded file is actually an image.
	contentType := header.Header.Get("Content-Type")

	switch contentType {
	case "image/jpeg", "image/png", "image/webp":
		// Allowed.
	default:
		http.Error(w, "Only JPEG, PNG, and WebP images are allowed", http.StatusBadRequest)
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

	key := fmt.Sprintf("profile-images/%d/avatar", userID)

	err = h.Storage.Upload(
		r.Context(),
		file,
		key,
		contentType,
	)
	if err != nil {
		http.Error(w, "Failed to upload profile image", http.StatusInternalServerError)
		return
	}

	err = h.UserRepository.UpdateProfileImageKey(userID, key)
	if err != nil {
		http.Error(w, "Failed to save profile image", http.StatusInternalServerError)
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
