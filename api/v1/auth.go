package v1

import (
	"github.com/gin-gonic/gin"
	"jester/pkg"
	"net/http"
)

// register godoc
// @Summary      Register a new user
// @Description  post data for user registration
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  Success
// @Failure		 400  {object}  Error
// @Router       /auth/register [post]
func register(c *gin.Context) {
	var input pkg.RegisterUser

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{E: err.Error()})
		return
	}

	err = pkg.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{E: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Success{Status: "success"})
}

// login godoc
// @Summary      Login a user
// @Description  post data to get a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  Success
// @Failure		 400  {object}  Error
// @Router       /auth/login [post]
func login(c *gin.Context) {
	var input pkg.LoginUser

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{E: err.Error()})
		return
	}

	token, err := pkg.Login(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{E: err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
