package server

import (
	"log"
	"net/http"
	"time"
)

// Config holds server configuration
type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Start starts the HTTP server with the given configuration
func Start(handler http.Handler, cfg Config) error {
	srv := &http.Server{
		Handler:      handler,
		Addr:         ":" + cfg.Port,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	log.Printf("Server starting on port %s", cfg.Port)
	return srv.ListenAndServe()
}

