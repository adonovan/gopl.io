# 연습문제 4.1

두 SHA256 해시에서 다른 비트의 개수를 세는 함수를 작성하라(2.6.2 절의 PopCount를 참조하라)

Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount from Section 2.6.2.)

<hr>

`PopCount`

```go
package popcount

// population count - 인구수 - 1로 지정된 비트 수
// pc[i]는 i의 인구수다.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount는 x의 인구수를 반환한다.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
```