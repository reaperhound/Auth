package controllers

import (
	"encoding/json"
	"errors"
	"jwt-auth/internal/models"
	"jwt-auth/internal/services"
	"net/http"
	"strings"
)

type AuthController struct {
	authService *services.AuthService
}

func NewUserController(service *services.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := userBodyValidation(user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	authErr := c.authService.SignUp(user.Username, user.Password)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Signed up"))
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	if err := userBodyValidation(user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	if err := c.authService.Login(user.Username, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Logged in"))
}

func userBodyValidation(user models.User) error {
	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return errors.New("username and password are required")
	}

	return nil
}
