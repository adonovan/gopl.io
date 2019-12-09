package echo13

import (
	"fmt"
	"strings"
	"testing"
)

func joinmethod(text []string) {
	fmt.Println(strings.Join(text, " "))
}

// START OMIT
func BenchmarkJoinmethod(b *testing.B) {
	for i := 0; i < b.N;  i++ {
		joinmethod([]string{"this", "is", "join", "text"})
	}
}
// END OMIT

// `go test -bench=.`
// this is join text
// 95224             10995 ns/op
// PASS
