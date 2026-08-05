package controllers

import (
	"jwt-auth/internal/services"
	"net/http"
)

type UserController struct {
	authService *services.AuthService
}

func NewUserController(service *services.AuthService) *UserController {
	return &UserController{
		authService: service,
	}
}

func (c *UserController) SignUp(w http.ResponseWriter, r *http.Request) {
	response := c.authService.SignUp()

	w.Write([]byte(response))
}
