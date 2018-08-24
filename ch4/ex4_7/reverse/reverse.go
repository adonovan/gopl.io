/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package reverse

// ex 4.7 Modify reverse the characters of a []byte slice that represents a
// utf-8 encoded string, in place. Can you do it without allocating new memory?

/*
// the original reverse function from ch4

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
*/

// Reverse reverses a slice of bytes in place.
func Reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

/*

$ go tool compile -m reverse/reverse.go
reverse/reverse.go:17:18: Reverse b does not escape

i, and j are local variables and are stack allocated.
b does not escape, so should be stack allocated, not heap allocated.
I cannot think of any way to implement this without at least allocating i, j
as local variables.

*/
