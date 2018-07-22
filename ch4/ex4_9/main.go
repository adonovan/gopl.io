/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// ex 4.9: Write a program wordfreq to report the frequency of each word in an
// input text file. Call input.Split(bufio.ScanWords) before the first call to
// Scan to break the input into words instead of lines.

package main

import (
	"log"
	"os"

	"github.com/guidorice/gopl.io/ch4/ex4_9/wordfreq"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: wordfreq filename\n")
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := wordfreq.Count(f)
	log.Printf("%v\n", data)
}
