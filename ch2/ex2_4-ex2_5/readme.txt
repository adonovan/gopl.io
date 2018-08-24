ex2.4 and ex2.5 notes

The popcount example already implements solutions and benchmarks for population
count (number of set bits) in a uint64. The lookup table implementation
(PopCount) is the fastest. It is surprising the BitCount function, Hacker's
Delight, Figure 5-2 is equally fast. But at what cost for readability?

gopl.io/ch2$ cd popcount/
gopl.io/ch2/popcount$ go test -v -bench=.

goos: darwin
goarch: amd64
pkg: github.com/guidorice/gopl.io/ch2/popcount
BenchmarkPopCount-4             	2000000000	         0.35 ns/op
BenchmarkBitCount-4             	2000000000	         0.35 ns/op
BenchmarkPopCountByClearing-4   	50000000	        27.7 ns/op
BenchmarkPopCountByShifting-4   	10000000	       138 ns/op
PASS
ok  	github.com/guidorice/gopl.io/ch2/popcount	4.403s
