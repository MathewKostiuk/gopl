// Package bank provides a concurrency-safe bank with one account.
package bank

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdrawal(amount int) bool {
	mu.Lock()
	defer mu.Unlock()

	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
func deposit(amount int) { balance += amount }
