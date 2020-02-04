package treesort

import (
	"testing"
)

func TestString(t *testing.T) {
	root := &tree{value: 3}
	root = add(root, 2)
	root = add(root, 4)
	if root.String() != "[2 3 4]" {
		t.Log(root)
		t.Fail()
	}
}
