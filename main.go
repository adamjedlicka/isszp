package main

import (
	"flag"

	"log"

	"encoding/json"

	"io/ioutil"

	"gitlab.fit.cvut.cz/isszp/isszp/cmd/install"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model/db"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server"
)

type Config struct {
	Server   server.Config
	Database db.Config
}

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal("Missing config file in ./config/ directory!")
	}

	cfg := Config{}
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		panic(err)
	}

	if flag.Arg(0) == "install" {
		install.InstallDatabase(cfg.Database)
		return
	}

	db.Configure(cfg.Database)
	db.Init()

	server.Configure(cfg.Server)
	server.Run()
}
