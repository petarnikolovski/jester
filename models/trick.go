package models

import (
	"gorm.io/gorm"
)

type Trick struct {
	gorm.Model
	Description string
	Instruction string
	SectionID   int
	Section     Section
}
