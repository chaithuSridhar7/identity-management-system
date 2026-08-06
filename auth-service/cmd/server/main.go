package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/database"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer db.Close()

	fmt.Println("Database connected successfully")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Auth service running")
	})

	fmt.Println("Server running on port 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
