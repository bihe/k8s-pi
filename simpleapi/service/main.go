package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var Version = "1.0.0"

// Payload defines a JSON structure which is returned by the API
type Payload struct {
	Version string `json:"version,omitempty"`
	Host    string `json:"hostname,omitempty"`
}

func getVersionHost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-type", "application/json")
		host, _ := os.Hostname()
		json.NewEncoder(w).Encode(Payload{Version: Version, Host: host})
	}
}

func getHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}

// --------------------------------------------------------------------------

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", getVersionHost())
	r.Get("/healthz", getHealth())

	fmt.Printf("🚀 up and running @ %s\n", fmt.Sprintf(":%s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
