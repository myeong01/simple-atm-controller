package cli

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/myeong01/simple-atm-controller/cmd/atm-controller/config/viper/bank"
	cardreadermock "github.com/myeong01/simple-atm-controller/pkg/cardreader/mock"
	cashbinmock "github.com/myeong01/simple-atm-controller/pkg/cashbin/mock"
	"testing"
	"unsafe"
)

var testController = &Controller{
	bankConfig: bank.Config{},
	cardReader: &cardreadermock.Controller{},
	cashBin:    &cashbinmock.Controller{},
}

func TestController_MainMenu(t *testing.T) {
	nextFn, _ := testController.MainMenu(args{})
	var expectedFn exec = testController.ReadCard

	// This does not allow you to verify that the function is really a method of testController
	// (methods of other instances have the same function address).
	nextFnPtrStr := fmt.Sprint(nextFn)
	expectedFnPtrStr := fmt.Sprint(expectedFn)
	if nextFnPtrStr != expectedFnPtrStr {
		t.Error("unexpected nextFn")
	}
}

func TestController_GoBackToMainMenu(t *testing.T) {
	// TODO implement test code
	t.Error("test code not implemented")
}

func TestController_ReadCard(t *testing.T) {
	// TODO implement test code
	t.Error("test code not implemented")
}

func TestController_GetBankFromCard(t *testing.T) {
	// TODO implement test code
	t.Error("test code not implemented")
}

func TestController_ValidateCard(t *testing.T) {
	// TODO implement test code
	t.Error("test code not implemented")
}

func TestController_SelectAccount(t *testing.T) {
	// TODO implement test code
	t.Error("test code not implemented")
}

func TestController_SelectActionForAccount(t *testing.T) {
	// TODO implement test code
	t.Error("test code not implemented")
}

func TestController_DoActionForAccount(t *testing.T) {
	// TODO implement test code
	t.Error("test code not implemented")
}

var equateFuncs = cmp.Comparer(func(x, y exec) bool {
	px := *(*unsafe.Pointer)(unsafe.Pointer(&x))
	py := *(*unsafe.Pointer)(unsafe.Pointer(&y))
	return px == py
})
