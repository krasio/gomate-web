package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/krasio/gomate"
	"github.com/soveran/redisurl"
)

func query(w http.ResponseWriter, r *http.Request) {
	kind := r.URL.Query().Get("kind")
	query := r.URL.Query().Get("q")

	// Connect to Redis
	redisUrl, ok := os.LookupEnv("GOMATE_REDIS_URL")
	if !ok {
		redisUrl = "redis://localhost:9999/0"
	}
	log.Printf("Using %s.\n", redisUrl)
	conn, err := redisurl.ConnectToURL(redisUrl)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	matches := gomate.Query(kind, query, conn)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}
