package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	User     string
	Password string
	Database string
}

var (
	cfg Config

	db *gorm.DB
	sq *sql.DB
)

func Configure(config Config) { cfg = config }

func Init() *gorm.DB {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s", cfg.User, cfg.Password, cfg.Database))
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

	return db
}
