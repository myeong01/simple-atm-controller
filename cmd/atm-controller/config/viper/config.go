package viper

import (
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/bank"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/cardreader"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/cashbin"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/controller"
	"github.com/spf13/viper"
)

type Config struct {
	CardReader cardreader.Config `yaml:"cardReader"`
	Bank       bank.Config       `yaml:"bank"`
	Controller controller.Config `yaml:"controller"`
	CashBin    cashbin.Config    `yaml:"cashBin"`
}

func SetConfigDefault(config *viper.Viper) {
	bank.SetConfigDefault("bank", config)
	cardreader.SetConfigDefault("cardReader", config)
	controller.SetConfigDefault("controller", config)
	cashbin.SetConfigDefault("cashBin", config)
}

func NewFromFile(filepath string) (*Config, error) {
	viperConfig := viper.New()

	SetConfigDefault(viperConfig)
	viperConfig.SetConfigFile(filepath)

	err := viperConfig.ReadInConfig()
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = viperConfig.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
