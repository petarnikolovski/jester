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

	result := db.Create(&s)
	if result.Error != nil {
		logger.Log.Error(result.Error)
		return &Section{}, result.Error
	}

	result = db.Preload("User").Preload("Level").First(&s, s.ID)
	if result.Error != nil {
		logger.Log.Error(result.Error)
		return &Section{}, result.Error
	}

	if s.SectionID != nil {
		parentSection := Section{}
		result = db.Model(&Section{}).First(&parentSection, s.SectionID)
		if result.Error != nil {
			logger.Log.Error(result.Error)
			return &Section{}, result.Error
		}

		if *s.UserID != *parentSection.UserID {
			return &Section{}, errors.New("User must be owner of the section")
		}

		result = db.Preload("Section").First(&s, s.ID)
	}

	return s, nil
}

type Level struct {
	gorm.Model
	Name string
}
