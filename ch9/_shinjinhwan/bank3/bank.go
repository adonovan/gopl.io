package bank

import "sync"

var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}


func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}