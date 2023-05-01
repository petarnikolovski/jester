package models

import (
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	Title       string
	Description string
	LevelID     uint
	Level       Level
	SectionID   *uint
	Section     *Section `gorm:"constraint:OnDelete:CASCADE;"`
}

type Level struct {
	gorm.Model
	Name string
}
