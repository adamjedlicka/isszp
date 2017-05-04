package server

import (
	"log"
	"net/http"
)

type Config struct {
	Address string
	Port    string
}

var (
	cfg Config
)

func Configure(config Config) { cfg = config }

func Run() {
	r := NewRouter()

	log.Println("Starting server at: ", cfg.Address+":"+cfg.Port)

	http.ListenAndServe(cfg.Address+":"+cfg.Port, r)
}
