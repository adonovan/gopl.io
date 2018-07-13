// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// ex 4.4: write a version of rotate that operates in a single pass.

package main

import (
	"fmt"

	"github.com/guidorice/gopl.io/ch4/ex4_4/rotate"
)

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(s)
	rotate.Rotate(s, 2)
	fmt.Println(s)
}
