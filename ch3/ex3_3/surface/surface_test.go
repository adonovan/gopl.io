/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package surface

import (
	"log"
	"testing"
)

var tests = [][]float64{
	// arg1, arg2
	{0, 0},
	{1, 1},
}

func TestCorner(t *testing.T) {
	for _, t := range tests {
		i, j, z, err := Corner(int(t[0]), int(t[1]))
		if err != nil {
			log.Fatalf("f: for args %v, got %v, %v, %v", t[0:2], i, j, z)
		}
	}
}

