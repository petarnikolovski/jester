package models

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"jester/databases/db"
	"strings"
)

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique;"`
	Password string `gorm:"not null;"`
}

func (u *User) Save() (*User, error) {
	result := db.GetDB().Create(&u)
	if result.Error != nil {
		log.Error(result.Error)
		return &User{}, result.Error
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}
