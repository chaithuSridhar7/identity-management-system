package main

import (
	"net/http"

	"github.com/chaithuSridhar7/identity-management-system/auth-service/internal/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HomeHandler)

	http.ListenAndServe(":8080", nil)
}
