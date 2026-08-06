package services

import (
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/models"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/repository"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/security"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (s *UserService) RegisterUser(
	username string,
	email string,
	password string,
) (*models.User, error) {

	hash, err := security.HashPassword(password)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: hash,
	}

	err = s.UserRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}