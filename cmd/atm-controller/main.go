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
	if err = controller.Run(); err != nil {
		log.Panicln("failed to run controller:", err)
	}
}
