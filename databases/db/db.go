package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var database *gorm.DB

func SetDB(db *gorm.DB) {
	database = db
}

func GetDB() *gorm.DB {
	if database == nil {
		log.Panic("Global database instance is not initialized")
	}
	return database
}
