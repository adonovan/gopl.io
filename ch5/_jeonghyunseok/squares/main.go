// square 프로그램은 상태를 가진 함수를 보여준다.package main

package main

import "fmt"

func main() {
	f := squares()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// go run main.go
// 신기하다!
