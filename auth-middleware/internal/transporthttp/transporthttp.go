package transporthttp

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/emmearn/gotlas.git/auth-middleware/internal/models"
)

type Middleware interface {
	Wrap(http.Handler) http.Handler
}

type API struct {
	authMiddleware Middleware
}

func NewAPI(authMiddleware Middleware) (*API, error) {
	if authMiddleware == nil {
		return nil, errors.New("error in http transport")
	}

	return &API{authMiddleware: authMiddleware}, nil
}

func (a *API) RegisterRoutes(router *mux.Router) {
	a.registerRoute(router, "GET", "/api/v1/whatever", handlerFunc)
}

func (a *API) registerRoute(router *mux.Router, method, path string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	router.Methods(method).Path(path).Handler(a.authMiddleware.Wrap(http.HandlerFunc(handlerFunc)))
}

func handlerFunc(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	res := &models.Response{
		Message: "Hi!",
	}
	b, _ := json.Marshal(res)
	w.Write(b)
}
