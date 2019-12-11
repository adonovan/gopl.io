package cc

import (
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	myslice := []string{"Penn", "Teller", "asdf", "asdf", "qwertz"}
	for i := 0; i < b.N; i++ {
		Concat(myslice)
	}
}

func BenchmarkJoin(b *testing.B) {
	myslice := []string{"Penn", "Teller", "asdf", "asdf", "qwertz"}
	for i := 0; i < b.N; i++ {
		Join(myslice)
	}
}
