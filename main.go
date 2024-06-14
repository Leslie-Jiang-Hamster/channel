package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())


	r.GET("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})

	r.NoRoute(func(c *gin.Context) {
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
	
	r.Run(":8000")
}