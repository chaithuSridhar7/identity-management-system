package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/models"
)

type fakeUserService struct {
	registerUser  *models.User
	registerError error
	loginUser     *models.User
	loginError    error
}

func (f *fakeUserService) RegisterUser(
	username string,
	email string,
	password string,
) (*models.User, error) {

	if f.registerError != nil {
		return nil, f.registerError
	}

	return f.registerUser, nil
}

func (f *fakeUserService) LoginUser(
	email string,
	password string,
) (*models.User, error) {

	if f.loginError != nil {
		return nil, f.loginError
	}

	return f.loginUser, nil
}

func TestRegisterHandler(t *testing.T) {

	fakeService := &fakeUserService{
		registerUser: &models.User{
			ID:       1,
			Username: "testuser",
			Email:    "test@example.com",
		},
	}

	handler := NewAuthHandler(fakeService)

	requestBody := RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to create request body: %v", err)
	}

	request := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer(body),
	)

	recorder := httptest.NewRecorder()

	handler.Register(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusCreated,
			recorder.Code,
		)
	}
}

func TestRegisterHandlerInvalidMethod(t *testing.T) {

	fakeService := &fakeUserService{}

	handler := NewAuthHandler(fakeService)

	request := httptest.NewRequest(
		http.MethodGet,
		"/register",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.Register(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusMethodNotAllowed,
			recorder.Code,
		)
	}
}

func TestRegisterHandlerInvalidJSON(t *testing.T) {

	fakeService := &fakeUserService{}

	handler := NewAuthHandler(fakeService)

	request := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBufferString("invalid json"),
	)

	recorder := httptest.NewRecorder()

	handler.Register(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}
}

func TestRegisterHandlerServiceError(t *testing.T) {

	fakeService := &fakeUserService{
		registerError: errors.New("registration failed"),
	}

	handler := NewAuthHandler(fakeService)

	requestBody := RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to create request body: %v", err)
	}

	request := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer(body),
	)

	recorder := httptest.NewRecorder()

	handler.Register(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}
}

func TestLoginHandler(t *testing.T) {

	fakeService := &fakeUserService{
		loginUser: &models.User{
			ID:       1,
			Username: "testuser",
			Email:    "test@example.com",
		},
	}

	handler := NewAuthHandler(fakeService)

	requestBody := LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to create request body: %v", err)
	}

	request := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBuffer(body),
	)

	recorder := httptest.NewRecorder()

	handler.Login(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusOK,
			recorder.Code,
		)
	}
}

func TestLoginHandlerInvalidMethod(t *testing.T) {

	fakeService := &fakeUserService{}

	handler := NewAuthHandler(fakeService)

	request := httptest.NewRequest(
		http.MethodGet,
		"/login",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.Login(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusMethodNotAllowed,
			recorder.Code,
		)
	}
}

func TestLoginHandlerInvalidJSON(t *testing.T) {

	fakeService := &fakeUserService{}

	handler := NewAuthHandler(fakeService)

	request := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBufferString("invalid json"),
	)

	recorder := httptest.NewRecorder()

	handler.Login(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}
}

func TestLoginHandlerInvalidCredentials(t *testing.T) {

	fakeService := &fakeUserService{
		loginError: errors.New("invalid email or password"),
	}

	handler := NewAuthHandler(fakeService)

	requestBody := LoginRequest{
		Email:    "wrong@example.com",
		Password: "wrongpassword",
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to create request body: %v", err)
	}

	request := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBuffer(body),
	)

	recorder := httptest.NewRecorder()

	handler.Login(recorder, request)

	if recorder.Code != http.StatusUnauthorized {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusUnauthorized,
			recorder.Code,
		)
	}
}
