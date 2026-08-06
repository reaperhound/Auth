package routes

import (
	"jwt-auth/internal/controllers"
	"net/http"
)

func RegisterProtectedRoutes(mux *http.ServeMux, c *controllers.ProtectedController) {
	mux.HandleFunc("GET /hey", c.Hey)
}
