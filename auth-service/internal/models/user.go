package models

import "time"

type User struct {
	ID              int       `json:"id"`
	Username        string    `json:"username"`
	DisplayName     string    `json:"display_name"`
	Email           string    `json:"email"`
	PasswordHash    string    `json:"-"`
	ProfileImageKey string    `json:"profile_image_key"`
	CreatedAt       time.Time `json:"created_at"`
}
