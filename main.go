package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"gitlab.fit.cvut.cz/isszp/isszp/cmd/install"
	"gitlab.fit.cvut.cz/isszp/isszp/src/controller"
	"gitlab.fit.cvut.cz/isszp/isszp/src/database"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model/db"
	"gitlab.fit.cvut.cz/isszp/isszp/src/server"
)

var dbUser string

type Config struct {
	Server     server.Config
	Database   database.Config
	Controller controller.Config
}

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&dbUser, "dbUser", "", "Prihlasovaci jmeno do databaze")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "install" {
		err := install.InstallConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal("Missing config file in ./config/ directory!")
	}

	cfg := Config{}
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		panic(err)
	}

	if dbUser != "" {
		cfg.Database.User = dbUser
	}

	if flag.Arg(0) == "install" {
		install.InstallDatabase(cfg.Database)
		return
	}

	controller.Configure(cfg.Controller)
	controller.Init()

	database.Configure(cfg.Database)
	gorm := database.Init()

	db.Init(gorm)

	server.Configure(cfg.Server)
	server.Run()
}
