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

type LoginResponse struct {
	Message string `json:"message"`
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

	accessToken, refreshToken, err := c.authService.Login(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send JSON
	w.Header().Set("Content-Type", "application/json")

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 15,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 24 * 7,
	})

	w.WriteHeader(http.StatusOK)

	response := LoginResponse{
		Message: "Logged in successfully",
	}

	json.NewEncoder(w).Encode(response)
}

func userBodyValidation(user models.User) error {
	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return errors.New("username and password are required")
	}

	return nil
}
