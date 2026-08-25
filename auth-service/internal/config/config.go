package config

import "os"

type Config struct {
	JWTSecret string
}

func Load() *Config {
	return &Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
