package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func tokenValid(c *gin.Context) error {
	bearerToken := c.Request.Header["Authorization"]
	if len(bearerToken) == 0 {
		return errors.New("Bearer token is missing.")
	}
	tokenString := strings.Split(bearerToken[0], " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return err
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return err
		} else {
			return err
		}
	}

	if token != nil && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
