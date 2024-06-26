package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine) {
	setUpV1Routes(router)
}

func setUpV1Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})
}