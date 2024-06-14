package main

import (
	"fmt"

	"github.com/channel/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	bootstrap.SetupRoutes(router)
	err := router.Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}
}