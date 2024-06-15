package main

import (
	"fmt"

	"github.com/channel/bootstrap"
	"github.com/channel/config"
	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.SetupEnv()
	env := config.GetConfig()
	router := gin.New()
	bootstrap.SetupRoutes(router)
	err := router.Run(fmt.Sprintf(":%d", env.AppConfig.Port))
	if err != nil {
		fmt.Println(err.Error())
	}
}