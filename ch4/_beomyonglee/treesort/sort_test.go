package treesort_test

import (
	"sort"
	"testing"

	"gopl.io/ch4/_beomyonglee/treesort"
)
import "math/rand"

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
