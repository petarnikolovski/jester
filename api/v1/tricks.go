package v1

import (
	"github.com/gin-gonic/gin"
	"jester/api"
	"jester/models"
	"jester/pkg"
	"net/http"
)

// createTrick godoc
// @Summary      Create a trick
// @Description  create a new trick and associate it with subsection
// @Tags         sections
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Trick
// @Failure		 400  {object}  Error
// @Router       /sections/{sectionId}/tricks [post]
func createTrick(c *gin.Context) {
	var input pkg.TrickInput
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

	parentID, err := api.StringToUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{E: "Invalid type for section ID"})
		return
	}

	trick, err := pkg.CreateTrick(input, user.(*models.User), parentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{E: err.Error()})
		return
	}

	c.JSON(http.StatusOK, trick)
}
