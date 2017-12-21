package main

import (
	"log"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/krasio/gomate-web/handlers"
	"github.com/krasio/gomate-web/version"
)

func main() {
	log.Printf(
		"GOMATE - commit: %s, build time: %s, release: %s.",
		version.Commit, version.BuildTime, version.Release,
	)

	port, ok := os.LookupEnv("GOMATE_PORT")
	if !ok {
		port = "8080"
	}

	router := handlers.Router()

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":" + port)
}
