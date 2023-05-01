package db

import (
	"errors"
	"gorm.io/gorm"
)

var database *gorm.DB

func SetDB(db *gorm.DB) {
	database = db
}

func GetDB() (*gorm.DB, error) {
	var err error
	if database == nil {
		err = errors.New("Global database instance is not initialized")
	}
	return database, err
}
