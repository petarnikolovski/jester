package pkg

import (
	"jester/databases/db"
	"jester/models"
)

type sectionLevel struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetLevels() ([]sectionLevel, error) {
	var levels []models.Level
	db, err := db.GetDB()
	if err != nil {
		return []sectionLevel{}, err
	}

	results := db.Find(&levels)

	var sectionLevels []sectionLevel
	for _, level := range levels {
		sectionLevels = append(sectionLevels, sectionLevel{ID: level.ID, Name: level.Name})
	}

	return sectionLevels, results.Error
}
