package security

import (
	"testing"
)

func TestGenerateAccessToken(t *testing.T) {
	secret := "test-secret"

	token, err := GenerateAccessToken(
		1,
		"test@example.com",
		secret,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if token == "" {
		t.Fatal("expected token, got empty string")
	}
}

func TestValidateAccessToken(t *testing.T) {
	secret := "test-secret"

	token, err := GenerateAccessToken(
		1,
		"test@example.com",
		secret,
	)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	parsedToken, err := ValidateAccessToken(token, secret)

	if err != nil {
		t.Fatalf("expected token to be valid, got %v", err)
	}

	if !parsedToken.Valid {
		t.Fatal("expected token to be valid")
	}
}

func TestValidateAccessTokenWrongSecret(t *testing.T) {
	token, err := GenerateAccessToken(
		1,
		"test@example.com",
		"correct-secret",
	)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	_, err = ValidateAccessToken(
		token,
		"wrong-secret",
	)

	if err == nil {
		t.Fatal("expected token validation to fail")
	}
}
