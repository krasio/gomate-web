package handlers

import (
	"net/http"
)

func readyz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
