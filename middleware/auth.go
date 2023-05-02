package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"jester/databases/db"
	"jester/logger"
	"jester/models"
	"jester/pkg"
	"net/http"
	"os"
	"strings"
)

func getToken(c *gin.Context) (*jwt.Token, error) {
	bearerToken := c.Request.Header["Authorization"]
	if len(bearerToken) == 0 {
		return nil, errors.New("Bearer token is missing.")
	}
	tokenString := strings.Split(bearerToken[0], " ")[1]

	token, err := jwt.ParseWithClaims(tokenString, &pkg.CustomJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})
	return token, err
}

func tokenValid(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}

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

func getUser(c *gin.Context) (*models.User, error) {
	token, err := getToken(c)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*pkg.CustomJWTClaims); ok && token.Valid {
		user := models.User{}
		db, err := db.GetDB()
		if err != nil {
			return nil, err
		}

		result := db.Model(&models.User{}).Where("email = ?", claims.Email).Take(&user)
		if result.Error != nil {
			return nil, result.Error
		}
		return &user, nil
	}
	return nil, errors.New("Token claims are not valid")
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenValid(c)
		if err != nil {
			logger.Log.Error(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user, err := getUser(c)
		if err != nil {
			logger.Log.Error(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user", user)

		c.Next()
	}
}
