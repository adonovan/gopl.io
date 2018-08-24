// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package rotate

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Rotate modifies in place a slice so elements rotated left n positions, in
// a single pass. See main() in ch4/rev/main.go.
func Rotate(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}
