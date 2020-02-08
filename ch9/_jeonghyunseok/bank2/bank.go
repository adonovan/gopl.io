// bank2

package bank2

var (
	sema = make(chan int, 1)
	balance int
)

func Deposit(amount int) {
	sema<-0
	balance += amount
	<-sema
}
func Balance() (b int ){
	sema<-0
	b = balance
	<-sema
	return
}