package middlewares

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const ClaimsContextKey contextKey = "claims"

type JWTVerifier interface {
	ParseAccessToken(tokenString string) (jwt.MapClaims, error)
}

func AuthMiddlerWare(jwtService JWTVerifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("access_token")
			if err != nil {
				if err == http.ErrNoCookie {
					http.Error(w, "Access Token not found", http.StatusUnauthorized)
					return
				}

				http.Error(w, "Failed to read Cookie", http.StatusInternalServerError)
				return
			}

			claims, err := jwtService.ParseAccessToken(cookie.Value)
			if err != nil {
				http.Error(w, "Invalid or expired access token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
