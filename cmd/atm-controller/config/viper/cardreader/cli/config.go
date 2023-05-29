package cli

import (
	"github.com/myeong01/simple-atm-controller/pkg/cardreader/cli"
	"github.com/spf13/viper"
)

type Config struct {
}

func SetConfigDefault(_ string, _ *viper.Viper) {
}

func (c *Config) GetController() (*cli.Controller, error) {
	return &cli.Controller{}, nil
}
