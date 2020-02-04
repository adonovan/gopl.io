package palindrome

import (
	"sort"
	"testing"
)

func TestPalindrome_oddNumElements(t *testing.T) {
	ints := []int{1, 2, 3, 4, 3, 2, 1}
	if !IsPalindrome(sort.IntSlice(ints)) {
		t.Fail()
	}
}

func TestPalindrome_evenNumElements(t *testing.T) {
	ints := []int{1, 2, 3, 3, 2, 1}
	if !IsPalindrome(sort.IntSlice(ints)) {
		t.Fail()
	}
}

func TestPalindrome_negative(t *testing.T) {
	ints := []int{1, 2, 3, 3, 2, 2}
	if IsPalindrome(sort.IntSlice(ints)) {
		t.Fail()
	}
}
