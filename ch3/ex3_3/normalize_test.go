/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package ex3_3

import (
	"log"
	"testing"
)

var tests = [][]float64{
	// min, max, value, expect
	{0, 10, 5, 0.5},
	{100, 300, 100, 0},
	{90, 1000, 1000, 1},
	{200, 400, 300, 0.5},
}

func TestNormalize(t *testing.T) {
	for _, t := range tests {
		value, _ := Normalize(t[0], t[1], t[2])
		if value != t[3] {
			log.Fatalf("Normalize: for args %v, %v, %v got %v, want %v",
				t[0], t[1], t[2], value, t[3])
		}
	}
}

