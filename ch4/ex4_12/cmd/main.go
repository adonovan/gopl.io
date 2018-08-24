/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"fmt"
	"os"

	"github.com/guidorice/gopl.io/ch4/ex4_12/xkcd"
)

const lastComicToken = 0

func main() {
	lastComic, err := xkcd.Get(lastComicToken)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("%v", lastComic)

	// TODO: from lastComic.Num to 1, fetch each json file, save to disk as single file metadata.json
	// TODO: add search interface
}
