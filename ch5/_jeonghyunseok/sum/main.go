// Sum 은 variadic function 을 보여준다.
// 즉, 파라미터 개수가 가변적인 함수이다.package main

package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(3, 4))
	fmt.Println(sum(1, 2, 3, 4))

	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))
}

func sum(v ...int) int {
	total := 0
	for _, v := range v {
		total += v
	}
	return total
}

// go run main.go
