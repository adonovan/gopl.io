// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */


// 4.3: Rewrite reverse to use an array pointer instead of a slice. n.b. in
// practice I cannot see why this would be any more useful than the slice
// version of reverse, because arrays are fixed-length.
package main

import (
	"fmt"
)

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println(a) // "[5 4 3 2 1 0]"
	reverse5(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
}

// reverse5 reverses a pointer to an array of 5 ints in place.
func reverse5(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
