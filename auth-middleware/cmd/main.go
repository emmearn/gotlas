package main

import (
	"log"
	"net/http"
	"time"

	auth "github.com/emmearn/gotlas.git/auth-middleware/internal/authentication"
	"github.com/emmearn/gotlas.git/auth-middleware/internal/transporthttp"

	"github.com/gorilla/mux"
)

func main() {
	authSvc, _ := auth.NewAuthService()
	middleware := auth.NewAuthMiddleware(*authSvc)

	r := mux.NewRouter()

	api, _ := transporthttp.NewAPI(middleware)
	api.RegisterRoutes(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
