package models

import (
	"gorm.io/gorm"
	"time"
)

type Trick struct {
	gorm.Model
	Description string
	Instruction string
	SectionID   int
	Section     Section
	CreatedAt   time.Time
}
