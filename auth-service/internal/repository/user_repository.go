package repository

import (
	"database/sql"

	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {

	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := r.DB.QueryRow(
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	return err
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {

	user := &models.User{}

	query := `
		SELECT id, username, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`

	err := r.DB.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindUserByID(id int) (*models.User, error) {

	user := &models.User{}

	var displayName sql.NullString
	var profileImageKey sql.NullString

	query := `
		SELECT id, username, display_name, email, password_hash, profile_image_key, created_at
		FROM users
		WHERE id = $1
	`

	err := r.DB.QueryRow(
		query,
		id,
	).Scan(
		&user.ID,
		&user.Username,
		&displayName,
		&user.Email,
		&user.PasswordHash,
		&profileImageKey,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	if displayName.Valid {
		user.DisplayName = displayName.String
	} else {
		user.DisplayName = user.Username
	}

	if profileImageKey.Valid {
		user.ProfileImageKey = profileImageKey.String
	}

	return user, nil
}

func (r *UserRepository) UpdateDisplayName(id int, displayName string) error {

	query := `
        UPDATE users
        SET display_name = $1
        WHERE id = $2
    `

	_, err := r.DB.Exec(
		query,
		displayName,
		id,
	)

	return err
}

func (r *UserRepository) UpdateProfileImageKey(id int, key string) error {
	query := `
		UPDATE users
		SET profile_image_key = $1
		WHERE id = $2
	`

	_, err := r.DB.Exec(
		query,
		key,
		id,
	)

	return err
}
