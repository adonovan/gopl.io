package main

import (
	"reflect"
	"testing"
)

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, a []byte, b []byte) {
	if reflect.DeepEqual(a, b) {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", string(a), reflect.TypeOf(a), string(b), reflect.TypeOf(b))
}

func BenchmarkUtf8Reverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Utf8Reverse([]byte("abcdefghijklmnopqrstuvwxyz가나다라마바사아자차카타파하❤"))
	}
}

func TestUtf8Reverse(t *testing.T) {
	b := []byte("⭐abc가나다라마바사아자차카타파하💗")
	Utf8Reverse(b)
	AssertEqual(t,
		b,
		[]byte("💗하파타카차자아사바마라다나가cba⭐"),
	)
}
