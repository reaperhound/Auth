package routes

import (
	"jwt-auth/internal/controllers"
	"jwt-auth/internal/middlewares"
	"jwt-auth/internal/services"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	fileSerivice := services.NewFileService()
	jwtService := services.NewJwtService()
	authService := services.NewAuthService(fileSerivice, jwtService)
	authController := controllers.NewUserController(authService)

	public := http.NewServeMux()
	RegisterAuthRoutes(public, authController)

	// Protected
	protected := http.NewServeMux()
	RegisterProtectedRoutes(protected)
	protectedWithAuth := middlewares.AuthMiddlerWare(jwtService)(protected)

	mux := http.NewServeMux()
	mux.Handle("/", public)
	mux.Handle("/private/", http.StripPrefix("/private", protectedWithAuth))

	return mux
}
