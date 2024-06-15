package bootstrap

import (
	"flag"
	"fmt"

	"github.com/channel/config"
)

func SetupEnv() error {
	path := flag.String("e", ".env", "Environment variables")
	if err := config.InitConfig(path); err != nil {
		return fmt.Errorf("Fail to load config: %v", err)
	}
	return nil
}