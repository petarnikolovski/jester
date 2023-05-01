package v1

import (
	"github.com/gin-gonic/gin"
	"jester/middleware"
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

	auth := v1.Group("/auth")

	{
		auth.POST("/register", register)
		auth.POST("/login", login)
	}

	section := v1.Group("/section")
	section.Use(middleware.AuthRequired())

	{
		section.GET("/levels", listLevels)
	}
}
