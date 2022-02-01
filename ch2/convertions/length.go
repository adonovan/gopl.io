// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package convertions

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%g meter", m) }
func (f Feet) String() string  { return fmt.Sprintf("%g feet", f) }

func Meter2Feet(m Meter) Feet {
	return Feet(m * 3.28084)
}

func Feet2Meter(f Feet) Meter {
	return Meter(f * 0.3048)
}

//!-
