package models

import (
	"gorm.io/gorm"
	"time"
)

type Section struct {
	gorm.Model
	Title       string
	Description string
	LevelID     uint
	Level       Level
	SectionID   *uint
	Section     *Section `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time
}

type Level struct {
	gorm.Model
	Name      string
	CreatedAt time.Time
}
