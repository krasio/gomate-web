package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/krasio/gomate-web/handlers"
	"github.com/krasio/gomate-web/version"
	"github.com/urfave/negroni"
)

func main() {
	log.Printf(
		"[GOMATE] commit: %s, build time: %s, release: %s.",
		version.Commit, version.BuildTime, version.Release,
	)

	port, ok := os.LookupEnv("GOMATE_PORT")
	if !ok {
		port = "8080"
	}

	router := handlers.Router()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	n := negroni.Classic()
	n.UseHandler(router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: n,
	}

	go func() {
		srv.ListenAndServe()
	}()
	log.Print("[GOMATE] Listening on :", port)

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("[GOMATE] Got SIGINT.")
	case syscall.SIGTERM:
		log.Print("[GOMATE] Got SIGTERM.")
	}

	log.Print("[GOMATE] Shutting down...")
	srv.Shutdown(context.Background())
	log.Print("[GOMATE] Done.")
}
