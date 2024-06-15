package bootstrap

import (
	"net/http"
	"strings"

	"github.com/channel/route"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	SetupMiddlewares(router)
	route.SetupAPIRoutes(router)
	SetupNoRoutes(router)
}

func SetupMiddlewares(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func SetupNoRoutes(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			c.String(http.StatusNotFound, "Not Found")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_msg": "Not Found",
			})
		}
	})
}