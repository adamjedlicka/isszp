package install

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gitlab.fit.cvut.cz/isszp/isszp/src/database"
)

func InstallDatabase(cfg database.Config) {
	log.Println("Installing database...")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/", cfg.User, cfg.Password))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	tx, err := db.Begin()

	_, err = tx.Exec("DROP DATABASE IF EXISTS `" + cfg.Database + "`")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	_, err = tx.Exec("CREATE DATABASE `" + cfg.Database + "`")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	_, err = tx.Exec("USE `" + cfg.Database + "`")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	{ // DDL SQL statements
		data, err := ioutil.ReadFile("config/ddl.sql")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		stmts := strings.Split(string(data), ";")
		stmts = stmts[:len(stmts)-1]

		for _, stmt := range stmts {
			_, err := tx.Exec(stmt)
			if err != nil {
				log.Println(err)
				fmt.Println(stmt)
				os.Exit(1)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Database creadetd succesfully...")
	log.Println("Filling database with initial data...")

	{ // DML SQL statements
		tx, err = db.Begin()

		data, err := ioutil.ReadFile("config/dml.sql")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		stmts := strings.Split(string(data), ";")
		stmts = stmts[:len(stmts)-1]

		for _, stmt := range stmts {
			_, err := tx.Exec(stmt)
			if err != nil {
				log.Println(err)
				fmt.Println(stmt)
				os.Exit(1)
			}
		}

		err = tx.Commit()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	log.Println("Database filled successfull..")
}
