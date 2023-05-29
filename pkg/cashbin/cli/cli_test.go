package cli

import (
	"github.com/myeong01/simple-atm-controller/pkg/testutils"
	"math/rand"
	"strconv"
	"testing"
)

func TestController_Deposit(t *testing.T) {
	var (
		testCase  = make([]int, 100)
		testInput = ""
	)
	for idx := range testCase {
		curRandomValue := rand.Int()
		testCase[idx] = curRandomValue
		testInput += strconv.Itoa(curRandomValue) + "\n"
	}

	deferFn, err := testutils.MockStdin(t, testInput)
	if err != nil {
		t.Fatal(err)
	}
	defer deferFn()

	c := Controller{}

	for _, value := range testCase {
		amount, err := c.Deposit()
		if err != nil {
			t.Error(err)
		} else {
			if amount != value {
				t.Error(amount, "!=", value)
			}
		}
	}
}

func TestController_Withdraw(t *testing.T) {
	// TODO implement test code
	t.Error("test code not implemented")
}

func TestController_IsAvailableToWithdraw(t *testing.T) {
	// TODO implement test code
	t.Error("test code not implemented")
}
