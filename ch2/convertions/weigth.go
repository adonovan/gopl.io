// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package convertions

import "fmt"

type Kilogram float64
type Pound float64

func (m Kilogram) String() string { return fmt.Sprintf("%g Kilograms", m) }
func (f Pound) String() string    { return fmt.Sprintf("%g Pounds", f) }

func Kilogram2Pound(m Kilogram) Pound {
	return Pound(m * 2.20462)
}

func Pound2Kilogram(f Pound) Kilogram {
	return Kilogram(f * 0.4536)
}

//!-
