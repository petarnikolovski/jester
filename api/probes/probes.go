package probes

import (
	"github.com/gin-gonic/gin"
	"jester/databases/db"
	"net/http"
)

func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"healthy": "OK"})
}

func readiness(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal server error"})
		return
	}

	result := db.Exec("SELECT 1")
	if result.Error != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Service unavailable"})
	}
	c.JSON(http.StatusOK, gin.H{"ready": "OK"})
}

func Routes(router *gin.Engine) {
	probe := router.Group("/")

	{
		probe.GET("/healthz", healthz)
		probe.GET("/ready", readiness)
	}
}
