package bootstrap

import (
	"flag"
	"fmt"

	"github.com/channel/env"
)

func SetupEnv() error {
	path := flag.String("e", ".env", "Environment variables")
	flag.Parse()
	if err := env.InitConfig(path); err != nil {
		return fmt.Errorf("Fail to load config: %v", err)
	}
	return nil
}