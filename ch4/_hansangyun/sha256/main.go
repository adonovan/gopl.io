package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	//%x : 배열이나 바이트 슬라이스의 모든 원소를 16진수로 출력
	//%t : boolean
	//%T : 값의 타입 표시

	zero(&c1)
	zero2(&c2)

	fmt.Printf("c1 = %x", c1)
	fmt.Printf("c2 = %x", c2)

}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	*ptr = [32]byte {}
}