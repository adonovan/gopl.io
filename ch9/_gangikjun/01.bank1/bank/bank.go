package bank

import "fmt"

var deposits = make(chan int) // deposit으로 amount를 보낸다
var balances = make(chan int) // balance를 받는다

// Deposit Deposit
func Deposit(amount int) { deposits <- amount }

// Balance Balance
func Balance() int { return <-balances }

func teller() {
	var balance int // balance는 teller 고루틴에 국한된다.
	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Println("balance += amount")
		case balances <- balance:
			fmt.Println("balance")
		}
	}
}

func init() {
	go teller()
}
