package mock

import (
	"github.com/myeong01/simple-atm-controller/pkg/bankapi/mock"
	"github.com/spf13/viper"
)

var (
	mockController *mock.Controller
)

func init() {
	mockController = &mock.Controller{}
}

type Config struct {
}

func SetConfigDefault(_ string, _ *viper.Viper) {
}

func (c *Config) GetController() (*mock.Controller, error) {
	return mockController, nil
}
