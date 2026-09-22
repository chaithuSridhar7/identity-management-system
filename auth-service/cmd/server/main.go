package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/database"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/handlers"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/middleware"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/repository"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/services"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer db.Close()

	fmt.Println("Database connected successfully")

	userRepository := repository.NewUserRepository(db)

	userService := services.NewUserService(userRepository, jwtSecret)

	authHandler := handlers.NewAuthHandler(userService)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	meHandler := handlers.NewMeHandler(userRepository)

	protectedMeHandler := middleware.AuthMiddleware(
		jwtSecret,
		meHandler,
	)

	http.Handle("/me", protectedMeHandler)

	router := http.DefaultServeMux

	handler := middleware.CORS(router)

	fmt.Println("Server running on port 8080")

	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
