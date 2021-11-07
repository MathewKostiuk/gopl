package bank

import (
	"fmt"
	"testing"
)

func TestBalance(t *testing.T) {
	Deposit(100)
	fmt.Println(Balance())
	fmt.Println(Withdrawal(100))
	fmt.Println(Balance())
}
