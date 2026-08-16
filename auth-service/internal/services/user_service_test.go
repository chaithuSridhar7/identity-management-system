package services

import (
	"errors"
	"testing"

	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/models"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/security"
)

type fakeUserRepository struct {
	createdUser       *models.User
	userToReturn      *models.User
	findUserError     error
	createUserError   error
}

func (f *fakeUserRepository) CreateUser(user *models.User) error {
	if f.createUserError != nil {
		return f.createUserError
	}

	f.createdUser = user
	return nil
}

func (f *fakeUserRepository) FindUserByEmail(email string) (*models.User, error) {
	if f.findUserError != nil {
		return nil, f.findUserError
	}

	return f.userToReturn, nil
}

func TestRegisterUser(t *testing.T) {
	repo := &fakeUserRepository{}

	service := NewUserService(repo)

	user, err := service.RegisterUser(
		"testuser",
		"test@example.com",
		"password123",
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user == nil {
		t.Fatal("expected user, got nil")
	}

	if user.Username != "testuser" {
		t.Errorf("expected username testuser, got %s", user.Username)
	}

	if user.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", user.Email)
	}

	if user.PasswordHash == "password123" {
		t.Error("password should not be stored as plaintext")
	}

	if !security.VerifyPassword("password123", user.PasswordHash) {
		t.Error("password hash does not match original password")
	}

	if repo.createdUser == nil {
		t.Fatal("expected CreateUser to be called")
	}
}

func TestRegisterUserRepositoryError(t *testing.T) {
	repo := &fakeUserRepository{
		createUserError: errors.New("database error"),
	}

	service := NewUserService(repo)

	user, err := service.RegisterUser(
		"testuser",
		"test@example.com",
		"password123",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if user != nil {
		t.Error("expected nil user when registration fails")
	}
}

func TestLoginUser(t *testing.T) {
	hash, err := security.HashPassword("password123")

	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	repo := &fakeUserRepository{
		userToReturn: &models.User{
			ID:           1,
			Username:     "testuser",
			Email:        "test@example.com",
			PasswordHash: hash,
		},
	}

	service := NewUserService(repo)

	user, err := service.LoginUser(
		"test@example.com",
		"password123",
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user == nil {
		t.Fatal("expected user, got nil")
	}

	if user.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", user.Email)
	}
}

func TestLoginUserWrongPassword(t *testing.T) {
	hash, err := security.HashPassword("password123")

	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	repo := &fakeUserRepository{
		userToReturn: &models.User{
			ID:           1,
			Username:     "testuser",
			Email:        "test@example.com",
			PasswordHash: hash,
		},
	}

	service := NewUserService(repo)

	user, err := service.LoginUser(
		"test@example.com",
		"wrongpassword",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if user != nil {
		t.Error("expected nil user, got user")
	}

	if err.Error() != "invalid email or password" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestLoginUserRepositoryError(t *testing.T) {
	repo := &fakeUserRepository{
		findUserError: errors.New("database error"),
	}

	service := NewUserService(repo)

	user, err := service.LoginUser(
		"test@example.com",
		"password123",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if user != nil {
		t.Error("expected nil user, got user")
	}

	if err.Error() != "invalid email or password" {
		t.Errorf("unexpected error: %v", err)
	}
}