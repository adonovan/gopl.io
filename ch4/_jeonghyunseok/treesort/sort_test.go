// Package treesort provides inserting sort using an unbalanced binary tree
package treesort_test

import (
	"math/rand"
	"sort"
	"testing"

	"gopl.io/ch4/_jeonghyunseok/treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

// go test . -v
// --- FAIL: TestSort (0.00s)
//     sort_test.go:19: not sorted: [47 10 12 19 7 2 21 20 16 15 30 32 7 33 39 3 9 23 43 44 21 38 0 22 9 28 37 11 6 23 31 40 41 18 23 15 36 24 37 34 49 16 48 8 20 37 1 21 1 10]
// FAIL
// FAIL    gopl.io/ch4/_jeonghyunseok/treesort     0.251s
// FAIL
