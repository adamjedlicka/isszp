package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	sq *sql.DB
)

func Init() {
	var err error
	db, err = gorm.Open("mysql", "root:@/isszp")
	if err != nil {
		panic(err)
	}

	sq = db.DB()

	err = sq.Ping()
	if err != nil {
		panic(err)
	}

	// Connect and check the server version
	var version string
	sq.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("Connected to:", version)
}

func NewUUID() string {
	var uuid string
	err := sq.QueryRow("SELECT UUID();").Scan(&uuid)
	if err != nil {
		panic(err)
	}

	return uuid
}
