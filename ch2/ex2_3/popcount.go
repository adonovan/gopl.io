// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// ex2.3 Rewrite PopCount to use a loop instead of a single expression. Compare
// the differences of the two versions. (see popcount_test.go for tests and
// benchmarks).

/*

$ go test -v -bench=.
=== RUN   TestPopCount
--- PASS: TestPopCount (0.00s)
=== RUN   TestPopCountDeadbeefs
--- PASS: TestPopCountDeadbeefs (0.00s)
=== RUN   TestPopCountCoffeeDeadbeef
--- PASS: TestPopCountCoffeeDeadbeef (0.00s)
=== RUN   TestPopCountWithLoop
--- PASS: TestPopCountWithLoop (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/guidorice/gopl.io/ch2/ex2_3
BenchmarkPopCount-4           	2000000000	         0.34 ns/op
BenchmarkPopCountWithLoop-4   	100000000	        22.2 ns/op
PASS
ok  	github.com/guidorice/gopl.io/ch2/ex2_3	2.969s

 */

package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
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

var bits = []int{0, 1, 2, 3, 4, 5, 6, 7}

// PopCountWithLoop returns the population count (number of set bits) of x.
// this looped implementation is ~65x slower than the original, above.
func PopCountWithLoop(x uint64) int {
	var count byte
	for i := range bits {
		count += pc[byte(x>>(uint(i)*8))]
	}
	return int(count)
}
