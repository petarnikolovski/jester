package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"jester/databases/db"
	"jester/models"
	"os"
	"strconv"
	"time"
)

type RegisterUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(data RegisterUser) error {
	user := models.User{}
	user.Email = data.Email
	user.Password = data.Password

	_, err := user.Save()
	return err
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

type CustomJWTClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func generateToken(user *models.User) (string, error) {
	tokenDuration, err := strconv.Atoi(os.Getenv("TOKEN_DURATION"))
	duration := time.Duration(tokenDuration)

	claims := CustomJWTClaims{
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	apiSecret := []byte(os.Getenv("API_SECRET"))
	tokenString, err := token.SignedString(apiSecret)
	return tokenString, err
}

func Login(data LoginUser) (string, error) {
	user := models.User{}
	db, err := db.GetDB()
	if err != nil {
		return "", err
	}

	if result := db.Model(&models.User{}).Where("email = ?", data.Email).Take(&user); result.Error != nil {
		return "", result.Error
	}

	err = verifyPassword(data.Password, user.Password)
	if err != nil {
		return "", err
	}

	token, err := generateToken(&user)
	if err != nil {
		return "", err
	}

	return token, nil
}
