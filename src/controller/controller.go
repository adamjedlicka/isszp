// Package controller is business logic of this application
// it servers as an middleground between view layer and model layer
package controller

import (
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

// Config is structure that holds configuration for controller package
type Config struct {
	Secret string
}

var (
	cfg Config
)

// Configure configures controller package with passed in config
func Configure(config Config) { cfg = config }

// Init initializes default state of controller package
// it is required to be called before initzializing view layer
func Init() {
	session.InitSessionStore(cfg.Secret)
}
