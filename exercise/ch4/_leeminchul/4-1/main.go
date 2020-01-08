package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
	"unsafe"
)

func CountSha256Diff(a, b [32]byte) (n int) {
	for i := 0; i < 32; i++ {
		n += bits.OnesCount8(a[i] ^ b[i])
	}
	return
}

func CountSha256DiffU64(a, b [32]byte) (n int) {
	// 8 byte * 4 배열을 xor
	// 배열의 크기가 8 byte 로 나누어 떨어지지 않으면 안됨 (unsafe.Pointer 사용)
	for i := 0; i < 32; i += 8 {
		n += bits.OnesCount64(*(*uint64)(unsafe.Pointer(&a[i])) ^ *(*uint64)(unsafe.Pointer(&b[i])))
	}
	return
}

func main() {
	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))

	fmt.Println(CountSha256Diff(s1, s2))
	fmt.Println(CountSha256DiffU64(s1, s2))
}
