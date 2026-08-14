package controllers

import (
	"encoding/json"
	"jwt-auth/internal/middlewares"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
)

type TwoFAController struct{}

func NewTwoFAController() *TwoFAController {
	return &TwoFAController{}
}

func (c *TwoFAController) Enroll2FA(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middlewares.ClaimsContextKey).(jwt.MapClaims)

	username, ok := claims["username"].(string)
	if !ok {
		http.Error(w, "invalid token claims", http.StatusUnauthorized)
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Go2FA",
		AccountName: username,
	})

	if err != nil {
		http.Error(w, "failed to generate secret", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"secret": key.Secret(),
		"qr_url": key.URL(),
	})
}
