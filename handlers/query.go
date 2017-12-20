package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/krasio/gomate"
	"github.com/soveran/redisurl"
)

func query(w http.ResponseWriter, r *http.Request) {
	kind := r.URL.Query().Get("kind")
	query := r.URL.Query().Get("q")

	// Connect to Redis
	log.Printf("Using %s.\n", "redis://localhost:9999/0")
	conn, err := redisurl.ConnectToURL("redis://localhost:9999/0")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	matches := gomate.Query(kind, query, conn)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}
