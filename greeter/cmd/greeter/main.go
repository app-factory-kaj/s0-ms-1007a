// Command greeter starts the greeter HTTP service.
//
// This is foundation-only scaffolding: it wires config, the HTTP server and
// the router bootstrap. The /hello handler itself lands in a follow-up issue
// on top of this layout, without restructuring it.
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
