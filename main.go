package main

import (
	"github.com/gin-gonic/gin"
	probes "jester/api/probes"
	"net/http"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func main() {
	router := gin.Default()

	probes.Routes(router)

	v1 := router.Group("/api/v1")

	{
		v1.GET("/hello", hello)
	}

	router.Run()
}
