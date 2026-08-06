package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		1,        // iterations
		64*1024,  // memory (64 MB)
		4,        // parallelism
		32,       // output length
	)

	return fmt.Sprintf(
		"%s:%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	), nil
}