// Package database describes common interface for comunicating with a database
// It uses GORM framework can default SQL querries can be used too
package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Config stores all config variables fro database package
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

// Configure configures database package with supplied config
func Configure(config Config) { cfg = config }

// Init connects to a database supplied in Config and returns its descriptor
// if any error occurs it panics. Must be called before controller or view Init
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
