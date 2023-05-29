package mock

type Controller struct {
}

func (c *Controller) Deposit() (int, error) {
	return 10, nil
}

func (c *Controller) Withdraw(amount int) error {
	return nil
}

func (c *Controller) IsAvailableToWithdraw(amount int) bool {
	return true
}
