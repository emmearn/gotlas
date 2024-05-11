package internal

import (
	"errors"
	"net/http"
)

type Middleware interface {
	Wrap(http.Handler) http.Handler
}

type API struct {
	authMiddleware Middleware
}

func NewAPI(authMiddleware Middleware) (*API, error) {
	if authMiddleware == nil {
		return nil, errors.New("ooooh")
	}

	return &API{authMiddleware: authMiddleware}, nil
}

func (a *API) RegisterRoutes(router *mux.Router) {
	// Al posto di `nil` ci andrebbe il tuo handler per la risorsa
	a.registerRoute(router, "GET", "/api/v1/whatever", nil)
}

func (a *API) registerRoute(router *mux.Router, method, path string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	router.Methods(method).Path(path).Handler(a.authMiddleware.Wrap(http.HandlerFunc(handlerFunc)))
}
