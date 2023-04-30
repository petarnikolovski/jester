package models

import (
	"gorm.io/gorm"
	"time"
)

type Section struct {
	gorm.Model
	Title       string
	Description string
	SectionID   uint
	Section     *Section `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time
}
