package main

import (
	"flag"

	"isszp/cmd/install"
	"isszp/src/model/db"
	"isszp/src/server"
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "install" {
		install.Init()
		return
	}

	db.Init()
	server.Run()
}
