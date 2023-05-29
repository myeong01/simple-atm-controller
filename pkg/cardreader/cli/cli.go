package cli

import "fmt"

type Controller struct {
}

func (c *Controller) ReadCardNumber() (string, error) {
	fmt.Print("Input Card Number : ")
	var cardNumber string
	fmt.Scan(&cardNumber)
	return cardNumber, nil
}
