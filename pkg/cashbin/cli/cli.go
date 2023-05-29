package cli

import "fmt"

type Controller struct {
}

func (c *Controller) Deposit() (int, error) {
	fmt.Print("Input : ")
	var amount int
	fmt.Scan(&amount)
	return amount, nil
}

func (c *Controller) Withdraw(amount int) error {
	fmt.Println("...giving ", amount, "$")
	return nil
}

func (c *Controller) IsAvailableToWithdraw(amount int) bool {
	return true
}
