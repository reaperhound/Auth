package routes

import (
	"jwt-auth/internal/controllers"
	"net/http"
)

func RegisterProtectedRoutes(mux *http.ServeMux) {
	pr := controllers.NewProtectedController()
	twoFa := controllers.NewTwoFAController()

	mux.HandleFunc("GET /hey", pr.Hey)
	mux.HandleFunc("GET /2fa/enroll", twoFa.Enroll2FA)
}
