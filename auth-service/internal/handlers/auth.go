package handlers

import (
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/services"
)

type AuthHandler struct {
	UserService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{
		UserService: userService,
	}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}