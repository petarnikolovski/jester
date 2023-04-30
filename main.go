package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jester/api/probes"
	"jester/models"
	"log"
	"net/http"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func main() {
	// dsn := "host=localhost user=user password=password dbname=jester port=5432 sslmode=disable TimeZone=UTC"
	dsn := "host=localhost user=user password=password dbname=jester port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&models.Section{}, &models.Trick{})

	router := gin.Default()

	probes.Routes(router)

	v1 := router.Group("/api/v1")

	{
		v1.GET("/hello", hello)
	}

	router.Run()
}
