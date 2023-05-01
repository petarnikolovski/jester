package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"jester/pkg"
	"net/http"
)

// listLevels godoc
// @Summary      List possible levels for sections
// @Description  get all levels
// @Tags         sections
// @Accept       json
// @Produce      json
// @Success      200  {object}  []pkg.sectionLevel
// @Router       /section/levels [get]
func listLevels(c *gin.Context) {
	levels, err := pkg.GetLevels()
	if err != nil {
		log.Panic(err)
	}

	c.JSON(http.StatusOK, levels)
}
