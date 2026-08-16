package services

import (
	"errors"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/models"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/security"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
}

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(repo UserRepository) *UserService {
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

func (s *UserService) LoginUser(
	email string,
	password string,
) (*models.User, error) {

	user, err := s.UserRepository.FindUserByEmail(email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !security.VerifyPassword(password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
