package pkg

import (
	"jester/databases/db"
	"jester/logger"
	"jester/models"
)

type SectionCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	LevelID     uint   `json:"levelID" binding:"required"`
	SectionID   *uint  `json:"sectionID"`
}

func CreateSection(data SectionCreate, user *models.User) (*models.Section, error) {
	section := models.Section{
		Title:       data.Title,
		Description: data.Description,
		UserID:      &user.ID,
		LevelID:     data.LevelID,
		SectionID:   data.SectionID,
	}

	s, err := section.Save()
	return s, err
}

func ListTopLevelSections(user *models.User) ([]models.Section, error) {
	db, err := db.GetDB()
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	var topLevelSections []models.Section
	results := db.Preload("User").Preload("Level").Where("user_id = ? AND section_id IS NULL", user.ID).Find(&topLevelSections)
	if results.Error != nil {
		logger.Log.Error(results.Error)
		return nil, results.Error
	}

	return topLevelSections, nil
}
