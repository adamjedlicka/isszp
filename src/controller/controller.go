package controller

import (
	"gitlab.fit.cvut.cz/isszp/isszp/src/server/session"
)

type Config struct {
	Secret string
}

var (
	cfg Config
)

func Configure(config Config) { cfg = config }

func Init() {
	session.InitSessionStore(cfg.Secret)
}
