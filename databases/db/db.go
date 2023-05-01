package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Panic("Global database instance is not initialized")
	}
	return DB
}

type Database interface {
	Connect() (*gorm.DB, error)
	InitDatabase()
}
