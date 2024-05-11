package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/emmearn/gotlas.git/auth-middleware/internal"

	"github.com/gorilla/mux"
)

func main() {
	authSvc, _ := internal.NewAuthService()
	middleware := internal.NewAuthMiddleware(*authSvc)

	r := mux.NewRouter()

	_, api := internal.NewAPI(middleware)
	api.RegisterRoutes(r, handler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	res := &internal.Response{
		Message: "Hi!",
	}
	b, _ := json.Marshal(res)
	w.Write(b)
}
