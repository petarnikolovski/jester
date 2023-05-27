package models

import (
	"errors"
	"gorm.io/gorm"
	"jester/databases/db"
	"jester/logger"
)

type Section struct {
	gorm.Model
	Title       string
	Description string
	LevelID     uint
	Level       Level
	SectionID   *uint
	Section     *Section `gorm:"constraint:OnDelete:CASCADE;"`
	UserID      *uint
	User        User `gorm:"constraint:OnDelete:CASCADE;"`
}

func (s *Section) Save() (*Section, error) {
	db, err := db.GetDB()
	if err != nil {
		return &Section{}, err
	}

	if s.LevelID == 2 && s.SectionID == nil {
		return &Section{}, errors.New("Subsection needs to have a parent section")
	}

	if result := db.Create(&s); result.Error != nil {
		logger.Log.Error(result.Error)
		return &Section{}, result.Error
	}

	if result := db.Preload("User").Preload("Level").First(&s, s.ID); result.Error != nil {
		logger.Log.Error(result.Error)
		return &Section{}, result.Error
	}

	if s.SectionID != nil {
		parentSection := Section{}

		if result := db.Model(&Section{}).First(&parentSection, s.SectionID); result.Error != nil {
			logger.Log.Error(result.Error)
			return &Section{}, result.Error
		}

		if *s.UserID != *parentSection.UserID {
			return &Section{}, errors.New("User must be owner of the section")
		}

		if result := db.Preload("Section").First(&s, s.ID); result.Error != nil {
			logger.Log.Error(result.Error)
			return &Section{}, result.Error
		}
	}

	return s, nil
}

type Level struct {
	gorm.Model
	Name string
}
