package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/database"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/handlers"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/repository"
	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/services"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer db.Close()

	fmt.Println("Database connected successfully")

	userRepository := repository.NewUserRepository(db)

	userService := services.NewUserService(userRepository)

	authHandler := handlers.NewAuthHandler(userService)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", authHandler.Register)	
	http.HandleFunc("/login", authHandler.Login)

	fmt.Println("Server running on port 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
