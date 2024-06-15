package env

import (
	"github.com/BurntSushi/toml"
)

var cfg IConfig

type IConfig struct {
	AppConfig IAppConfig
}

type IAppConfig struct {
	Port int
}

func InitConfig(path *string) error {
	_, err := toml.DecodeFile(*path, &cfg)
	return err
}

func GetConfig() IConfig {
	return cfg
}