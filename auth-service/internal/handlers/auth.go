package handlers

import (
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/models"
)

type UserService interface {
	RegisterUser(username string, email string, password string) (*models.User, error)
	LoginUser(email string, password string) (*models.User, error)
}

type AuthHandler struct {
	UserService UserService
}

func NewAuthHandler(userService UserService) *AuthHandler {
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