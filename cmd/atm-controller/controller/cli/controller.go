package cli

import (
	"fmt"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/bank"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/controller"
	"github.com/myeong01/simple-atm-controller/pkg/account"
	"github.com/myeong01/simple-atm-controller/pkg/bankapi"
	"github.com/myeong01/simple-atm-controller/pkg/cardidentifier"
	"github.com/myeong01/simple-atm-controller/pkg/cardreader"
	"github.com/myeong01/simple-atm-controller/pkg/cashbin"
)

type Controller struct {
	bankConfig bank.Config
	cardReader cardreader.Interface
	cashBin    cashbin.Interface
}

type args struct {
	err        error
	msg        string
	cardNumber string
	bankApi    bankapi.Interface
	account    account.Interface
}

type exec func(a args) (exec, args)

func (c *Controller) Run() error {
	var (
		curExec  exec = c.MainMenu
		nextArgs args = args{}
	)
	curExec, nextArgs = curExec(nextArgs)
	for ; curExec != nil && nextArgs.err == nil; curExec, nextArgs = curExec(nextArgs) {
	}
	return nextArgs.err
}

func (c *Controller) SetBankConfig(bankConfig bank.Config) controller.Interface {
	c.bankConfig = bankConfig
	return c
}

func (c *Controller) SetCardReader(cardReader cardreader.Interface) controller.Interface {
	c.cardReader = cardReader
	return c
}

func (c *Controller) SetCashBin(cashBin cashbin.Interface) controller.Interface {
	c.cashBin = cashBin
	return c
}

func (c *Controller) MainMenu(_ args) (exec, args) {
	fmt.Println("Welcome to Bear Robotics Bank")
	fmt.Println("Insert your card for the next step")
	return c.ReadCard, args{}
}

func (c *Controller) GoBackToMainMenu(arg args) (exec, args) {
	if arg.msg != "" {
		fmt.Println(arg.msg)
	}
	fmt.Println("Go back to Main Menu")
	return c.MainMenu, args{}
}

func (c *Controller) ReadCard(_ args) (exec, args) {
	cardNumber, err := c.cardReader.ReadCardNumber()
	if err != nil {
		return c.GoBackToMainMenu, args{msg: "failed to read card: " + err.Error()}
	}
	return c.GetBankFromCard, args{cardNumber: cardNumber}
}

func (c *Controller) GetBankFromCard(arg args) (exec, args) {
	bankType, err := cardidentifier.GetBankTypeFromCardNumber(arg.cardNumber)
	if err != nil {
		return c.GoBackToMainMenu, args{msg: "invalid card number: cannot find bank: " + err.Error()}
	}
	bankApi, err := c.bankConfig.GetController(bankType)
	if err != nil {
		return c.GoBackToMainMenu, args{msg: "[" + bankType + "] not yet supported"}
	}
	arg.bankApi = bankApi
	return c.ValidateCard, arg
}

func (c *Controller) ValidateCard(arg args) (exec, args) {
	fmt.Print("Input PIN : ")
	var pinNumber string
	fmt.Scan(&pinNumber)
	isValid, err := arg.bankApi.IsValidPin(arg.cardNumber, pinNumber)
	if !isValid {
		msg := "invalid pin"
		if err != nil {
			msg += ": " + err.Error()
		}
		return c.GoBackToMainMenu, args{msg: msg}
	}
	return c.SelectAccount, arg
}

func (c *Controller) SelectAccount(arg args) (exec, args) {
	accounts, err := arg.bankApi.ListAccountsByCardNumber(arg.cardNumber)
	if err != nil {
		return c.GoBackToMainMenu, args{msg: "failed to list accounts"}
	}
	fmt.Println("0 : Go back to menu")
	for i, ac := range accounts {
		fmt.Println(i+1, ":", ac.Name())
	}
	fmt.Print("Account Number : ")
	var accountNumber uint32
	fmt.Scan(&accountNumber)
	if len(accounts) < int(accountNumber) {
		fmt.Println("Invalid Account Number")
		return c.SelectAccount, arg
	}
	if accountNumber == 0 {
		return c.GoBackToMainMenu, args{}
	}
	arg.account = accounts[accountNumber-1]
	return c.SelectActionForAccount, arg
}

func (c *Controller) SelectActionForAccount(arg args) (exec, args) {
	fmt.Println("0 : Go back to menu")
	for i, ac := range bankapi.ActionTypes {
		fmt.Println(i+1, ":", ac)
	}
	fmt.Print("Action Number : ")
	var actionNumber uint32
	fmt.Scan(&actionNumber)
	if len(bankapi.ActionTypes) < int(actionNumber) {
		fmt.Println("Invalid Action Number")
		return c.SelectActionForAccount, arg
	}
	if actionNumber == 0 {
		return c.GoBackToMainMenu, args{}
	}
	arg.msg = bankapi.ActionTypes[actionNumber-1]
	return c.DoActionForAccount, arg
}

func (c *Controller) DoActionForAccount(arg args) (exec, args) {
	var msg string
	switch arg.msg {
	case bankapi.ActionTypeSeeBalance:
		amount, err := arg.bankApi.SeeBalance(arg.cardNumber, arg.account.Id())
		if err != nil {
			msg = "failed to see balance: " + err.Error()
		} else {
			msg = fmt.Sprint("Balance :", amount)
		}
	case bankapi.ActionTypeDeposit:
		fmt.Println("put money to the cash bin")
		amount, err := c.cashBin.Deposit()
		if err != nil {
			msg = "failed to get money"
		} else {
			if err := arg.bankApi.Deposit(arg.cardNumber, arg.account.Id(), amount); err != nil {
				msg = "failed to deposit: " + err.Error()
			} else {
				msg = "Success deposit"
			}
		}
	case bankapi.ActionTypeWithdraw:
		fmt.Println("Enter the withdraw amount")
		var amount int
		fmt.Scan(&amount)
		if !c.cashBin.IsAvailableToWithdraw(amount) {
			msg = "cash bin has insufficient cash reserves"
		} else {
			if err := arg.bankApi.Withdraw(arg.cardNumber, arg.account.Id(), amount); err != nil {
				msg = "failed to withdraw: " + err.Error()
			} else {
				if err := c.cashBin.Withdraw(amount); err != nil {
					msg = "failed to get money from cash bin: " + err.Error()
				}
			}
		}
	default:
		msg = "unknown action type [" + arg.msg + "]"
	}
	return c.GoBackToMainMenu, args{msg: msg}
}
