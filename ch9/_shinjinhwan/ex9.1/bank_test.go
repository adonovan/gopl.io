package bank_test

import (
	"../ex9.1"
	"fmt"
	"testing"
	"time"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank.Deposit(100)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		time.Sleep(time.Second * 1) // waiting for deposit
		bank.Withdraw(100)
		fmt.Println("Withdraw , =", bank.Balance())
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done

	if got, want := bank.Balance(), 200; got != want {
		t.Errorf("Balance = ${got}, want #{want}")
	}
}