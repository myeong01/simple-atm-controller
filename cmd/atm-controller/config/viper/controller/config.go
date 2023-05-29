package controller

import (
	"errors"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/controller/cli"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/controller"
	"github.com/spf13/viper"
)

const (
	TypeCli = "cli"
)

type Config struct {
	Type string     `yaml:"type"`
	Cli  cli.Config `yaml:"cli"`
}

func SetConfigDefault(basePath string, config *viper.Viper) {
	config.SetDefault(basePath+".type", TypeCli)
	cli.SetConfigDefault(basePath+".cli", config)
}

func (c *Config) GetController() (controller.Interface, error) {
	switch c.Type {
	case TypeCli:
		return c.Cli.GetController()
	}
	return nil, errors.New("unknown controller type")
}
