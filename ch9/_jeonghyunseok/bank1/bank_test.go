package bank

import (
	"testing"
	"fmt"
)

func TestBank(t *testing.T) {
	done := make(chan int)

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- 0
	}()

	// Bob
	go func() {
		Deposit(100)
		done<- 1
	}()

	<-done
	<-done

	if got, want := Balance(), 300; got!=want {
		t.Errorf("Balance: %d, want %d", got, want)
	}
}

// go test .