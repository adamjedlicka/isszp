package main

import (
	"flag"

	"gitlab.fit.cvut.cz/isszp/isszp/cmd/install"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model/db"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server"
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
