package echo13

import (
	"fmt"
	"strings"
	"testing"
)

var testText = []string{"this", "is", "join", "text"}

func joinmethod(text []string) {
	fmt.Println(strings.Join(text, " "))
}

// START OMIT
func BenchmarkJoinmethod(b *testing.B) {
	for i := 0; i < b.N;  i++ {
		joinmethod(testText)
	}
}
// END OMIT

// `go test -bench=.`
// this is join text
// 95224             10995 ns/op
// PASS
