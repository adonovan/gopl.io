/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package wordfreq

import (
	"bufio"
	"io"
)

// Counts is represents word counts in a map[string]int.
type Counts map[string]int

// Count scans the in reader, splits into words, and returns a Counts struct.
func Count(in io.Reader) Counts {
	c := Counts{}
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		c[word]++
	}
	return c
}
