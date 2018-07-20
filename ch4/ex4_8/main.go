// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// ex 4.8 Modify charcount to count letters, digits, and so on in their unicode
// categories, using functions like unicode.isLetter.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	ex4_8 "github.com/guidorice/gopl.io/ch4/ex4_8/charcount"
)

func main() {
	buf := bytes.Buffer{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}
	res := ex4_8.Charcount(buf.String())
	fmt.Printf("%+v\n", res)
}
