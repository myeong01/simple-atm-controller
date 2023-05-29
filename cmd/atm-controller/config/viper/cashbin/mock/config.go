package mock

import (
	"github.com/myeong01/simple-atm-controller/pkg/cashbin/mock"
	"github.com/spf13/viper"
)

type Config struct {
}

func SetConfigDefault(_ string, _ *viper.Viper) {
}

func (c *Config) GetController() (*mock.Controller, error) {
	return &mock.Controller{}, nil
}
