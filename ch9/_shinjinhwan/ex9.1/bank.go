package bank

import "fmt"

// bank1 vㅡ로그램에 Withdraw(amount int) bool 함수를 추하가라.
// 결과는 거래 성공 또는 자금 부족으로 인한 거래 실패를 표시해야 한다.
// 관리 고루틴으로 전송하는 메시지는 출금액 및 관리 고루틴이 Withdraw 불리언 결과를 반환할 새 채널 두가지가 있어야 한다.

var deposits = make(chan int)
var withdraws = make(chan int)
var confirms = make(chan bool)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }

func Withdraw(amount int) bool {
	withdraws <- amount
	confirm := <-confirms
	fmt.Println(confirm)
	if confirm != true {
		fmt.Println("Transaction Failed")
	} else {
		fmt.Println("Transaction Sucess")
	}

	return confirm
}

func Balance() int 		 { return <-balances }

func teller() {
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraws:
			if amount > balance {
				confirms <- false
			} else {
				balance -= amount
				confirms <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}