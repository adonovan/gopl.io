package main

import "fmt"

// arr 를 k 번 왼쪽으로 rotate
func rotateJuggling(slice []int, k int) {
	length := len(slice)
	JumpIndex := func(i int) int {
		if length <= i {
			i -= length
		}
		return i
	}

	gcdResult := gcd(k, length)
	for i := 0; i < gcdResult; i++ {
		temp := slice[i]
		var nextIndex int
		for j := i; true; {
			nextIndex = JumpIndex(j + k)
			if nextIndex == i {
				slice[j] = temp
				break
			}
			slice[j] = slice[nextIndex]
			j = nextIndex
		}
	}
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotateJuggling(a[:], 2)
	fmt.Println(a) // [2 3 4 5 6 7 8 9 0 1]
}
