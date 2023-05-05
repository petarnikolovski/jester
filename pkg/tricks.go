package pkg

import "jester/models"

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
