package models

import (
	"errors"
	"gorm.io/gorm"
	"jester/databases/db"
	"jester/logger"
)

type Trick struct {
	gorm.Model
	Description string
	Instruction string
	SectionID   *uint
	Section     Section `gorm:"constraint:OnDelete:CASCADE;"`
	UserID      *uint
	User        User `gorm:"constraint:OnDelete:CASCADE;"`
}

func (t *Trick) Save() (*Trick, error) {
	db, err := db.GetDB()
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	section := Section{}
	result := db.Where("id = ?", t.SectionID).First(&section)
	if result.Error != nil {
		logger.Log.Error(result.Error)
		return nil, result.Error
	}
	if *section.UserID != *t.UserID {
		return nil, errors.New("User does not own the section")
	}

	level := Level{}
	db.Where("name = ?", "Subsection").First(&level)
	if level.ID != section.LevelID {
		return nil, errors.New("Tricks can be added to level 2 sections only")
	}

	result = db.Create(&t)
	if result.Error != nil {
		logger.Log.Error(result.Error)
		return nil, result.Error
	}

	db.Preload("User").Preload("Section").First(&t)

	return t, nil
}
