package db

import "github.com/jinzhu/gorm"

var db *gorm.DB

func Init(database *gorm.DB) { db = database }

func NewUUID() string {
	var uuid string
	err := db.DB().QueryRow("SELECT UUID();").Scan(&uuid)
	if err != nil {
		panic(err)
	}

	return uuid
}
