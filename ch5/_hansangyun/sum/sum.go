package main

import (
	"fmt"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}
func main() {
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(sum(1,3,5,7,9))

	values := []int{1, 2, 3, 4,}
	fmt.Println(sum(values...)) // 슬라이스와 가변 인자 타입은 서로 다름
}
