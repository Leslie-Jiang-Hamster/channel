package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})
}