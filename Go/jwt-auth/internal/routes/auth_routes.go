package routes

import (
	"jwt-auth/internal/controllers"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux, c *controllers.AuthController) {
	mux.HandleFunc("POST /signup", c.SignUp)
	mux.HandleFunc("POST /login", c.Login)
	mux.HandleFunc("POST /refresh", c.Refresh)
}
