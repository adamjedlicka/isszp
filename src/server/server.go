package server

import (
	"log"
	"net/http"
)

var (
	address = "localhost"
	port    = "8080"
)

func Run() {
	r := NewRouter()

	log.Println("Starting server at: ", address+":"+port)

	http.ListenAndServe(address+":"+port, r)
}
