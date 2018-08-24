/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

/* ex 3.13 write const declarations for KB, MB up through YB as compactly as you
can. use powers of 10. note that iota has no exponentation operator. */

// FIXME: ZB and YB overflows int at runtime. instead of defining consts,
// better to use big.Int and big.Exp.

package main

import (
	"fmt"
	"strconv"

	"github.com/guidorice/gopl.io/ch3/ex3_11/comma"
)

// bytes in powers of 10
const (
	k = 1000
	KB
	MB = k * k
	GB = MB * k
	TB = GB * k
	PB = TB * k
	EB = PB * k
	ZB = EB * k // const ok, but overflows int at runtime
	YB = ZB * k // const ok, but overflows int at runtime
)

func main() {
	labels := []string{"KB", "MB", "GB", "TB", "PB", "EB"}
	powers := []int{KB, MB, GB, TB, PB, EB}
	for i, n := range powers {
		label := labels[i]
		formatted := comma.Comma(strconv.Itoa(n))
		fmt.Printf("%s = %s\n", label, formatted)
	}
}
