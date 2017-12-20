package main

import (
	"os"

	"github.com/codegangsta/negroni"
	"github.com/krasio/gomate-web/handlers"
)

func main() {
	port, ok := os.LookupEnv("GOMATE_PORT")
	if !ok {
		port = "8080"
	}

	router := handlers.Router()

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":" + port)
}
