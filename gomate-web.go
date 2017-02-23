package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krasio/gomate"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/soveran/redisurl"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	kind := r.URL.Query().Get("kind")
	query := r.URL.Query().Get("q")

	// Connect to Redis
	fmt.Printf("Using %s.\n", "redis://localhost:9999/0")
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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", queryHandler).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
