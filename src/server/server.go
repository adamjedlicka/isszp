// Package server is a representation of HTTP server and view layer
package server

import (
	"log"
	"net/http"
)

// Config containes all configuration for the view layer
type Config struct {
	Address string
	Port    string
}

var (
	cfg Config
)

// Configure configures model layer with supplied config
func Configure(config Config) { cfg = config }

// Run inits and runs the HTTP server. Has to be run as last (after model & controller layer)
func Run() {
	r := NewRouter()

	log.Println("Starting server at: ", cfg.Address+":"+cfg.Port)

	http.ListenAndServe(cfg.Address+":"+cfg.Port, r)
}
