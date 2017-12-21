package handlers

import (
	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", query).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz)

	return r
}
