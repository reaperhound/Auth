package routes

import (
	"jwt-auth/internal/controllers"
	"jwt-auth/internal/services"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	fileSerivice := services.NewFileService()
	jwtService := services.NewJwtService()
	authService := services.NewAuthService(fileSerivice, jwtService)
	authController := controllers.NewUserController(authService)

	mux := http.NewServeMux()
	RegisterAuthRoutes(mux, authController)

	return mux
}
