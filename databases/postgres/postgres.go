package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jester/models"
	"log"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	var err error
	// dsn := "host=localhost user=user password=password dbname=jester port=5432 sslmode=disable TimeZone=UTC"
	dsn := "host=localhost user=user password=password dbname=jester port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return DB, err
}

func seedLevels(db *gorm.DB) error {
	var levels []models.Level
	results := db.Find(&levels)

	if results.Error != nil {
		log.Panic(results.Error)
	}

	if results.RowsAffected == 0 {
		section := models.Level{Name: "Section"}
		result := db.Create(&section)
		if result.Error != nil {
			log.Panic(result.Error)
		}

		subsection := models.Level{Name: "Subsection"}
		result = db.Create(&subsection)
		if result.Error != nil {
			log.Panic(result.Error)
		}
	}

	return nil
}

func InitDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.Section{}, &models.Level{}, &models.Trick{})
	if err != nil {
		log.Panic(err)
	}

	err = seedLevels(db)
	if err != nil {
		log.Panic(err)
	}
}
