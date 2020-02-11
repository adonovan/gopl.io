package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.RWMutex
	balance int
)

// Deposit Deposit
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance = balance + amount
}

// Balance Balance
func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func main() {
	// Deposit [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		if i%100 == 0 {
			fmt.Println(Balance())
		}
		go func(amount int) {
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	fmt.Println(Balance(), (1000+1)*1000/2)

	if got, want := Balance(), (1000+1)*1000/2; got != want {
		fmt.Printf("Balance = %d, want %d", got, want)
	}
}
