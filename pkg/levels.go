package pkg

import (
	"gorm.io/gorm"
	"jester/models"
)

type sectionLevel struct {
	ID   uint
	Name string
}

func GetLevels(db *gorm.DB) ([]sectionLevel, error) {
	var levels []models.Level
	results := db.Find(&levels)

	var sectionLevels []sectionLevel
	for _, level := range levels {
		sectionLevels = append(sectionLevels, sectionLevel{ID: level.ID, Name: level.Name})
	}

	return sectionLevels, results.Error
}
