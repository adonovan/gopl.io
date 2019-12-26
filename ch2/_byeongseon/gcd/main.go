// 최대 공약수
// 공약수 중 가장 큰 공약수

package main

import (
	"fmt"
)

func main() {
	a, b := 12, 18
	fmt.Printf("%d와 %d의 최대 공약수는 %d이다.\n", a, b, gcd(a, b))
}

func gcd(x, y int) int {
	// [유클리드 호제법]
	// 1. 만약 a < b 이면, a와 b를 바꾼다.
	// 2. a, b에 대해서 a를 b로 나눈 나머지를 r을 구한다.
	// 2. a와 b의 최대공약수는 b와 r의 최대공약수와 같다.
	for y != 0 {
		x, y = y, x%y
		fmt.Printf("%d, %d\n", x, y)
	}
	return x
}
