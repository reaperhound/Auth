package controllers

import (
	"encoding/json"
	"jwt-auth/internal/models"
	"jwt-auth/internal/services"
	"net/http"
	"strings"
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
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		http.Error(w, "username and password are required", http.StatusBadRequest)
		return
	}

	authErr := c.authService.SignUp(user.Username, user.Password)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Signed up"))
}
