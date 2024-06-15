package main

import (
	"fmt"

	"github.com/channel/bootstrap"
	"github.com/channel/env"
	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.SetupEnv()
	env := env.GetConfig()
	router := gin.New()
	bootstrap.SetupRoutes(router)
	err := router.Run(fmt.Sprintf(":%d", env.AppConfig.Port))
	if err != nil {
		fmt.Println(err.Error())
	}
}