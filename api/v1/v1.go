package v1

import (
	"github.com/gin-gonic/gin"
)

type Success struct {
	Status string `json:"status"`
}

type Error struct {
	E string `json:"error"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	section := v1.Group("/section")

	{
		section.GET("/levels", listLevels)
	}

	auth := v1.Group("/auth")

	{
		auth.POST("/register", register)
		auth.POST("/login", login)
	}
}
