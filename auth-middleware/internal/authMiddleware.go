package internal

import (
	"context"
	"encoding/json"
	"net/http"
)

type AuthMiddleware struct {
	// dipendenze da chiamare per authnz
}

func NewAuthMiddleware(
// dipendenze
) *AuthMiddleware {
	return &AuthMiddleware{
		// dipendenze
	}
}

type Response struct {
	Message string `json:"message"`
}

// Wrap soddisfa l'interfaccia dichiarata in transporthttp.go
func (a AuthMiddleware) Wrap(next http.Handler) httl.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// fai quello che ti serve per authnz
		// questo è un esempio di come estrarre gli header
		token := r.Header.Get("Authorization")

		// Esempio: chiami la tua dep e ti restituisce merda
		authResponse, err := a.authService.ValidateToken(ctx, token)
		if err != nil {
			// qui controlli l'errore, ma per semplicità restituiamo sempre 400
			res := Response{
				Message: "oooooooh",
			}

			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(res); err != nil {
				// In caso qualcosa sia andato male con l'encoding
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		// inserisci eventuali dati nel contesto se vuoi
		ctx = context.WithValue(ctx, "token", token)
		ctx = context.WithValue(ctx, "user", user)

		// chiami il prossimo handler nella catena
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
