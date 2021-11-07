// Package bank provides a concurrency-safe bank with one account.
package bank

type transaction struct {
	amount   int
	approved chan bool
}

var deposits = make(chan int)            // send amount to deposit
var withdrawals = make(chan transaction) // send amount to withdrawal
var balances = make(chan int)            // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdrawal(amount int) bool {
	var approved = make(chan bool)
	trx := transaction{amount, approved}
	withdrawals <- trx
	return <-approved
}

func teller() {
	var balance int //  balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case trx := <-withdrawals:
			if balance-trx.amount >= 0 {
				balance -= trx.amount
				trx.approved <- true
			} else {
				trx.approved <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
