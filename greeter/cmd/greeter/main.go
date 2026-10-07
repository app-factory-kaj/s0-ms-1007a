// Command greeter starts the greeter HTTP service.
//
// It wires config, the HTTP server and the router, which serves the
// component's GET /hello contract (see internal/server).
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"greeter/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	mux := server.NewRouter()

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("greeter listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
