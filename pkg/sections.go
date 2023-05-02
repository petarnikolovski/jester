package pkg

import (
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
