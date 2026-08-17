package routes

import (
	"jwt-auth/internal/controllers"
	"jwt-auth/internal/services"
	"net/http"
)

func RegisterProtectedRoutes(mux *http.ServeMux) {
	// Dependencies
	qrService := services.NewQRService()
	fileService := services.NewFileService()

	// Services
	twoFaService := services.NewTwoFASerive(qrService, fileService)

	// Controllers
	pr := controllers.NewProtectedController()
	twoFa := controllers.NewTwoFAController(twoFaService)

	mux.HandleFunc("GET /hey", pr.Hey)
	mux.HandleFunc("GET /2fa/enroll", twoFa.Enroll2FA)
}
