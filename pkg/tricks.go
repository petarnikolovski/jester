package pkg

import (
	"jester/databases/db"
	"jester/logger"
	"jester/models"
)

type TrickInput struct {
	Description string `json:"description"`
	Instruction string `json:"instruction" binding:"required"`
}

func CreateTrick(data TrickInput, user *models.User, sectionID *uint) (*models.Trick, error) {
	trick := models.Trick{
		Description: data.Description,
		Instruction: data.Instruction,
		SectionID:   sectionID,
		UserID:      &user.ID,
	}

	t, err := trick.Save()
	return t, err
}

func ListTricks(user *models.User, sectionID *uint) ([]models.Trick, error) {
	db, err := db.GetDB()
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	tricks := []models.Trick{}
	result := db.Where("section_id = ? AND user_id = ?", sectionID, user.ID).Find(&tricks)
	if result.Error != nil {
		logger.Log.Error(err)
		return nil, err
	}

	db.Preload("User").Preload("Section").Find(&tricks)

	return tricks, nil
}
