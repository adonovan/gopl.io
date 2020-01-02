// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

type tree struct {
	value       int
	left, right *tree
}

// Sort는 값을 직접 정렬한다.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues는 t의 원소를 values에 순서대로 추가하고
// 결과 슬라이스를 반환한다.
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
