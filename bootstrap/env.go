package bootstrap

import (
	"flag"
	"fmt"

	"github.com/channel/env"
)

func SetupEnv() {
	path := flag.String("e", ".env", "Environment variables")
	flag.Parse()
	if err := env.LoadEnv(path); err != nil {
		fmt.Println(fmt.Errorf("Fail to load config: %v", err))
	}
}