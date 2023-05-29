package cli

import (
	"github.com/myeong01/simple-atm-controller/pkg/testutils"
	"math/rand"
	"strconv"
	"testing"
)

func TestController_ReadCardNumber(t *testing.T) {
	var (
		testCase  = make([]string, 100)
		testInput = ""
	)
	for idx := range testCase {
		curRandomValue := strconv.Itoa(rand.Int())
		testCase[idx] = curRandomValue
		testInput += curRandomValue + "\n"
	}

	deferFn, err := testutils.MockStdin(t, testInput)
	if err != nil {
		t.Fatal(err)
	}
	defer deferFn()

	c := Controller{}

	for _, value := range testCase {
		cardNumber, err := c.ReadCardNumber()
		if err != nil {
			t.Error(err)
		} else {
			if cardNumber != value {
				t.Error(cardNumber, "!=", value)
			}
		}
	}
}
