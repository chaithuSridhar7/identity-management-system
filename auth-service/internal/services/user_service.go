package services

import (
	"errors"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/models"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/security"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id int) (*models.User, error)
}

type UserService struct {
	UserRepository UserRepository
	JWTSecret      string
}

func NewUserService(repo UserRepository, jwtSecret string) *UserService {
	return &UserService{
		UserRepository: repo,
		JWTSecret:      jwtSecret,
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
) (*models.LoginResponse, error) {

	user, err := s.UserRepository.FindUserByEmail(email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !security.VerifyPassword(password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	token, err := security.GenerateAccessToken(
		user.ID,
		user.Email,
		s.JWTSecret,
	)

	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}
