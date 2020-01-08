# 연습문제 4.3

reverse<sup>*</sup> 를 재작성해 슬라이스 대신 배열 포인터를 사용하라

Rewrite reverse to use an array pointer instead of a slice.

<hr>

`rev.go`

```go
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
```