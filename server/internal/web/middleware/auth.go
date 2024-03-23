package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/R894/bugtracker/internal/web/response"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	ContextKeyUsername contextKey = "username"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)

		if token == "" {
			response.Error(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		valid, username := verifyToken(token)
		if !valid {
			response.Error(w, http.StatusUnauthorized, "Invalid Token")
			return
		}
		ctx := context.WithValue(r.Context(), ContextKeyUsername, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetContextUsername(ctx context.Context) (string, error) {
	username := ctx.Value(ContextKeyUsername)
	if username == nil {
		return "", errors.New("username not found in context")
	}
	return username.(string), nil
}

func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}

	return parts[1]
}

func verifyToken(tokenString string) (bool, string) {
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, ""
	}

	if !token.Valid {
		return false, ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, ""
	}
	return true, claims["username"].(string)
}

func CreateToken(username string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("Error creating token:", err)
		return "", err
	}

	return tokenString, nil
}

func ApplyAuthMiddleware(r chi.Router) {
	r.Use(AuthMiddleware)
}
