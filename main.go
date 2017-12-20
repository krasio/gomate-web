package main

import (
	"github.com/codegangsta/negroni"
	"github.com/krasio/gomate-web/handlers"
)

func main() {
	router := handlers.Router()

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
