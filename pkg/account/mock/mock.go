package mock

import (
	"github.com/myeong01/simple-atm-controller/pkg/bankapi"
)

type Account struct {
	bankApi    bankapi.Interface
	name       string
	id         string
	cardNumber string
}

func New(bankApi bankapi.Interface, name, id, cardNumber string) *Account {
	return &Account{bankApi: bankApi, name: name, id: id, cardNumber: cardNumber}
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) Id() string {
	return a.id
}
