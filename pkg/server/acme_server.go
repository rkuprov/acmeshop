package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// AcmeServer wraps standard http.Server and allows for additional functionality.
type AcmeServer struct {
	Server *http.Server
}

// Begin registers the routes and starts the server
func (a *AcmeServer) Begin() {
	r := mux.NewRouter()
	registerRoutes(r)
	s := http.Server{
		Handler: r,
		Addr:    "localhost:8080",
	}
	s.ListenAndServe()
}
