package v1

import (
	"github.com/gin-gonic/gin"
	"jester/models"
	"jester/pkg"
	"net/http"
)

// createSection godoc
// @Summary      Create a section or subsection
// @Description  create a new section or subsection
// @Tags         sections
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Section
// @Failure		 400  {object}  Error
// @Router       /auth/register [post]
func createSection(c *gin.Context) {
	var input pkg.SectionCreate

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{E: err.Error()})
		return
	}

	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, Error{E: "User not found"})
		return
	}

	section, err := pkg.CreateSection(input, user.(*models.User))
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{E: err.Error()})
		return
	}

	c.JSON(http.StatusOK, section)
}
