package cashbin

type Interface interface {
	Deposit() (int, error)
	Withdraw(amount int) error
	IsAvailableToWithdraw(amount int) bool
}
