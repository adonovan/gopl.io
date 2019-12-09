package echo13

import (
	"fmt"
	"testing"
)

func convmethod(text []string) {
	var s, sep  string

	for i := 1; i < len(text); i++ {
		s += sep + text[i]
		sep = " "
	}

	fmt.Println(s)
}

func BenchmarkConvMethod(b *testing.B) {
	for i := 0; i < b.N;  i++ {
		convmethod([]string{"this", "is", "conv", "text"})
	}
}

// `go test -bench=.`
// is conv text
// 120285             10334 ns/op
// PASS