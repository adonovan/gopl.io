/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package rotate

import (
	"reflect"
	"testing"
)


func TestRotate(t *testing.T) {
	s := []int{1,2}
	expect := []int{2,1}
	n := 1
	Rotate(s, n)
	if !reflect.DeepEqual(s, expect) {
		t.Errorf("Rotate: got %v, expect %v", s, expect)
	}

	s = []int{1,2,3,4,5}
	expect = []int{3,4,5,1,2}
	n = 2
	Rotate(s, n)
	if !reflect.DeepEqual(s, expect) {
		t.Errorf("Rotate: got %v, expect %v", s, expect)
	}
}
