package controller

import (
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/bank"
	"github.com/myeong01/simple-atm-controller/pkg/cardreader"
	"github.com/myeong01/simple-atm-controller/pkg/cashbin"
)

type Interface interface {
	Run() error
	SetBankConfig(bankConfig bank.Config) Interface
	SetCardReader(cardReader cardreader.Interface) Interface
	SetCashBin(cashBin cashbin.Interface) Interface
}
