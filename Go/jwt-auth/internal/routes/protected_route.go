package routes

import (
	"jwt-auth/internal/controllers"
	"jwt-auth/internal/services"
	"net/http"
)

func RegisterProtectedRoutes(mux *http.ServeMux) {
	twoFaService := services.NewTwoFASerive()

	pr := controllers.NewProtectedController()
	twoFa := controllers.NewTwoFAController(twoFaService)

	mux.HandleFunc("GET /hey", pr.Hey)
	mux.HandleFunc("GET /2fa/enroll", twoFa.Enroll2FA)
}
