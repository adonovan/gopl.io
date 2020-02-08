package bank3

import "sync"
var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}


func Balance() (b int) {
	mu.Lock()
	b = balance
	mu.Unlock()
	return
}