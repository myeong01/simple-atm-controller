package mock

import (
	"github.com/myeong01/simple-atm-controller/pkg/account"
	"github.com/myeong01/simple-atm-controller/pkg/account/mock"
)

type Controller struct {
}

func (c *Controller) IsValidPin(_, _ string) (bool, error) {
	return true, nil
}

func (c *Controller) ListAccountsByCardNumber(cardNumber string) ([]account.Interface, error) {
	return []account.Interface{
		mock.New(c, "name1", "001", cardNumber),
		mock.New(c, "name2", "002", cardNumber),
		mock.New(c, "name3", "003", cardNumber),
	}, nil
}

func (c *Controller) SeeBalance(cardNumber, accountId string) (int, error) {
	return 0, nil
}

func (c *Controller) Deposit(cardNumber, accountId string, amount int) error {
	return nil
}

func (c *Controller) Withdraw(cardNumber, accountId string, amount int) error {
	return nil
}
