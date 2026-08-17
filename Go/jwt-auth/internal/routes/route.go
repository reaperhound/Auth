package routes

import (
	"jwt-auth/internal/controllers"
	"jwt-auth/internal/middlewares"
	"jwt-auth/internal/services"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	// Dependencies
	fileService := services.NewFileService()
	jwtService := services.NewJwtService()
	qrService := services.NewQRService()

	// Services
	authService := services.NewAuthService(fileService, jwtService)
	twoFaService := services.NewTwoFASerive(qrService, fileService)

	// Controller
	authController := controllers.NewUserController(authService)
	prController := controllers.NewProtectedController()
	twoFaController := controllers.NewTwoFAController(twoFaService)

	public := http.NewServeMux()
	RegisterAuthRoutes(public, authController)

	protected := http.NewServeMux()
	RegisterProtectedRoutes(protected, prController, twoFaController)
	protectedWithAuth := middlewares.AuthMiddlerWare(jwtService)(protected)

	mux := http.NewServeMux()
	mux.Handle("/", public)
	mux.Handle("/private/", http.StripPrefix("/private", protectedWithAuth))

	return mux
}
