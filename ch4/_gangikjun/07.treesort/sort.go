package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

func main() {
	data := make([]int, 5)
	for i := range data {
		data[i] = rand.Int() % 50
	}

	fmt.Println(data)

	Sort(data)
	if !sort.IntsAreSorted(data) {
		log.Fatalf("not sorted: %v", data)
	}

	fmt.Println(data)
}

type tree struct {
	value       int
	left, right *tree
}

// Sort 는 값을 직접 정렬한다
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// &tree{value: value} 반환과 같다.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
