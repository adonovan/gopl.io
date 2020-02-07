// package bank is concurency-safe bank
// with one account

package bank

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int){
	deposits <- amount
}
func Balance() int {
	return <-balances
}

func teller() {
	var balance int 
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances<- balance:
		}
	}
}

func init() {
	go teller()
}

/*
1) As soon as package "bank" is imported, init() will be invoked
2) Then, go teller() 
	- It keeps the value balance, but it only be changed and fetched by the chan
	- for loop keep choose
		1) <-deposits comes, then balance will be increased
		2) balances get read, then balance is written to that channel 
3) Deposit() and Balance() is exposed function. 


*/
