/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package ex3_3

import "fmt"

// Normalize returns value in range [0, 1] from the domain range [min, max].
func Normalize(min, max, value float64) (float64, error) {
	if value < min {
		return 0, fmt.Errorf("error: Normalize: value < domain minimum")
	}
	if value > max {
		return 0, fmt.Errorf("error: Normalize: value > domain max")
	}
	return (value - min) / (max - min), nil
}

