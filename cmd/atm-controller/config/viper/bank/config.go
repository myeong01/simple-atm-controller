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
	Mock mock.Config `yaml:"mock"`
}

func SetConfigDefault(basePath string, config *viper.Viper) {
	mock.SetConfigDefault(basePath+".mock", config)
}

func (c *Config) GetController(bankType string) (bankapi.Interface, error) {
	switch bankType {
	case TypeMock:
		return c.Mock.GetController()
	}
	return nil, errors.New("unknown bank type")
}
