package bank

import (
	"errors"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/bank/mock"
	"github.com/myeong01/simple-atm-controller/pkg/bankapi"
	"github.com/spf13/viper"
)

const (
	TypeMock = "mock"
)

type Config struct {
	Type string      `yaml:"type"`
	Mock mock.Config `yaml:"mock"`
}

func SetConfigDefault(basePath string, config *viper.Viper) {
	config.SetDefault(basePath+".type", TypeMock)
	mock.SetConfigDefault(basePath+".mock", config)
}

func (c *Config) GetController() (bankapi.Interface, error) {
	switch c.Type {
	case TypeMock:
		return c.Mock.GetController()
	}
	return nil, errors.New("unknown bank type")
}
