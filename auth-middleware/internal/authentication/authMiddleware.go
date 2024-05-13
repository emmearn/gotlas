package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/emmearn/gotlas.git/auth-middleware/internal/models"
)

type AuthMiddleware struct {
	authService AuthService
}

func NewAuthMiddleware(authService AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// Wrap soddisfa l'interfaccia dichiarata in transporthttp.go
func (a AuthMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := r.Header.Get("Authorization")
		user := r.Header.Get("User")

		_, err := a.authService.ValidateToken(ctx, token)
		if err != nil {
			res := models.Response{
				Message: "error in middleware",
			}

			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(res); err != nil {
				// In caso qualcosa sia andato male con l'encoding
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		ctx = context.WithValue(ctx, "token", token)
		ctx = context.WithValue(ctx, "user", user)

		// prossimo handler nella catena
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
