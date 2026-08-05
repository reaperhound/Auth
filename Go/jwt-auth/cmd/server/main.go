package main

import (
	"jwt-auth/internal/controllers"
	"jwt-auth/internal/services"
	"log"
	"net/http"
)

func main() {
	fileSerivice := services.NewFileService()
	authService := services.NewAuthService(fileSerivice)

	authController := controllers.NewUserController(authService)

	http.HandleFunc("/signup", authController.SignUp)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
