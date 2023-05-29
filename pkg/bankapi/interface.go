package bankapi

import "github.com/myeong01/simple-atm-controller/pkg/account"

const (
	ActionTypeSeeBalance = "See Balance"
	ActionTypeDeposit    = "Deposit"
	ActionTypeWithdraw   = "Withdraw"
)

var ActionTypes = []string{
	ActionTypeSeeBalance,
	ActionTypeDeposit,
	ActionTypeWithdraw,
}

type Interface interface {
	IsValidPin(cardNumber, pinNumber string) (bool, error)
	ListAccountsByCardNumber(cardNumber string) ([]account.Interface, error)
	SeeBalance(cardNumber, accountId string) (int, error)
	Deposit(cardNumber, accountId string, amount int) error
	Withdraw(cardNumber, accountId string, amount int) error
}
