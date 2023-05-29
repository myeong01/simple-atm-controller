package cardreader

import (
	"errors"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/cardreader/cli"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/cardreader/mock"
	"github.com/myeong01/simple-atm-controller/pkg/cardreader"
	"github.com/spf13/viper"
)

const (
	TypeMock = "mock"
	TypeCli  = "cli"
)

type Config struct {
	Type string      `yaml:"type"`
	Mock mock.Config `yaml:"mock"`
	Cli  cli.Config  `yaml:"cli"`
}

func SetConfigDefault(basePath string, config *viper.Viper) {
	config.SetDefault(basePath+".type", TypeCli)
	mock.SetConfigDefault(basePath+".mock", config)
	cli.SetConfigDefault(basePath+".cli", config)
}

func (c *Config) GetController() (cardreader.Interface, error) {
	switch c.Type {
	case TypeMock:
		return c.Mock.GetController()
	case TypeCli:
		return c.Cli.GetController()
	}
	return nil, errors.New("unknown cardReader type")
}
