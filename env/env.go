package env

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var env IEnv

type IEnv struct {
	AppEnv IAppEnv
	DBEnv IDBEnv
}

type IAppEnv struct {
	Port int
}

type IDBEnv struct {
	Connection string
	Host string
	Port int
	Database string
	Username string
	Password string
	Charset string
	Max_open_conns int
	Max_idle_conns int
	Max_life_seconds int
}

func LoadEnv(path *string) error {
	_, err := toml.DecodeFile(*path, &env)
	fmt.Printf("LoadEnv: %v", env)
	return err
}

func GetEnv() IEnv {
	return env
}