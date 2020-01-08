// Package treesort provides inserting sort using an unbalanced binary tree
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root) // 직접 변경
}

// tree to []int
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
		return &tree{value: value}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	s := []int{5, 1, 6, 7, 9, 2}
	Sort(s)
	fmt.Println(s)
}
