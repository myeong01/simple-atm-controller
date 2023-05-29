package main

import (
	"flag"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper"
	"log"
)

func main() {
	configFilePath := flag.String("config", "", "Define config file path")
	flag.Parse()
	if configFilePath == nil {
		log.Panicln("config option must be defined")
		return
	}
	config, err := viper.NewFromFile(*configFilePath)
	if err != nil {
		log.Panicln("failed to read config:", err)
	}
	controller, err := config.Controller.GetController()
	if err != nil {
		log.Panicln("failed to get controller:", err)
	}
	cardReader, err := config.CardReader.GetController()
	if err != nil {
		log.Panicln("failed to get card reader:", err)
	}
	cashBin, err := config.CashBin.GetController()
	if err != nil {
		log.Panicln("failed to get cashBin:", err)
	}
	controller = controller.
		SetBankConfig(config.Bank).
		SetCardReader(cardReader).
		SetCashBin(cashBin)
	if err = controller.Run(); err != nil {
		log.Panicln("failed to run controller:", err)
	}
}
