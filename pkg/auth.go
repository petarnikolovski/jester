package pkg

import (
	"jester/models"
)

type RegisterUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(data RegisterUser) error {
	user := models.User{}
	user.Email = data.Email
	user.Password = data.Password

	_, err := user.Save()
	return err
}
