package controllers

import (
	"html/template"
	"jwt-auth/internal/middlewares"
	"jwt-auth/internal/services"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type TwoFAController struct {
	twoFAService *services.TwoFAService
}

func NewTwoFAController(twoFaService *services.TwoFAService) *TwoFAController {
	return &TwoFAController{
		twoFAService: twoFaService,
	}
}

func (c *TwoFAController) Enroll2FA(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middlewares.ClaimsContextKey).(jwt.MapClaims)

	username, ok := claims["username"].(string)
	if !ok {
		http.Error(w, "invalid token claims", http.StatusInternalServerError)
		return
	}

	data, err := c.twoFAService.GenerateEnrollment(username)
	if err != nil {
		http.Error(w, "failed to generate 2FA", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/2fa.html")
	if err != nil {
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}
}
