package v1

import (
	"github.com/gin-gonic/gin"
	"jester/databases/postgres"
	"jester/pkg"
	"log"
	"net/http"
)

func listLevels(c *gin.Context) {
	levels, err := pkg.GetLevels(postgres.DB)
	if err != nil {
		log.Panic(err)
	}

	c.JSON(http.StatusOK, levels)
}
