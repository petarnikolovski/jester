package v1

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/section/levels", listLevels)
	}
}
