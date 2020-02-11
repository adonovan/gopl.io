package main

var (
	sema    = make(chan struct{}, 1)
	balance int
)

// Deposit Deposit
func Deposit(amount int) {
	sema <- struct{}{} // 토큰 획득
	balance = balance + amount
	<-sema // 토큰 해제
}

// Balance Balance
func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

func main() {

}
