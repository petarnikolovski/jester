package probes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"healthy": "OK"})
}

func readiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ready": "OK"})
}

func Routes(router *gin.Engine) {
	probe := router.Group("/")

	{
		probe.GET("/healthz", healthz)
		probe.GET("/ready", readiness)
	}
}
