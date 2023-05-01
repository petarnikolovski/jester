package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jester/logger"
	"jester/models"
	"os"
)

func Connect() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_POSTGRES_SSLMODE")

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, name, port, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Panic(err)
	}
	return db, err
}

func seedLevels(db *gorm.DB) error {
	var levels []models.Level
	results := db.Find(&levels)

	if results.Error != nil {
		logger.Log.Panic(results.Error)
	}

	if results.RowsAffected == 0 {
		section := models.Level{Name: "Section"}
		result := db.Create(&section)
		if result.Error != nil {
			logger.Log.Panic(result.Error)
		}

		subsection := models.Level{Name: "Subsection"}
		result = db.Create(&subsection)
		if result.Error != nil {
			logger.Log.Panic(result.Error)
		}
	}

	return nil
}

func InitDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Section{}, &models.Level{}, &models.Trick{})
	if err != nil {
		logger.Log.Panic(err)
	}

	err = seedLevels(db)
	if err != nil {
		logger.Log.Panic(err)
	}
}
