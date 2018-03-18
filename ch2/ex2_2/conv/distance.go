/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package conv

// distance conversions

import (
	"errors"
	"flag"
	"fmt"
	"math"
)

type Feet float64
type Meters float64
type Yards float64

const (
	FeetPerMeter  = 3.28084
	YardsPerMeter = 1.09361
	FeetPerYard   = 3
	ftFlag = "ft"
	mFlag= "m"
	yFlag = "y"
)

var (
	ft, m, y *float64
)

// Stringer interface
func (f Feet) String() string   { return fmt.Sprintf("%g feet", f) }
func (m Meters) String() string { return fmt.Sprintf("%g meters", m) }
func (y Yards) String() string  { return fmt.Sprintf("%g yards", y) }

// Feet.Meters() -> Meters
func (f Feet) Meters() Meters {
	return Meters(f / FeetPerMeter)
}

// Yards.Meters() -> Meters
func (y Yards) Meters() Meters {
	return Meters(y / YardsPerMeter)
}

// Meters.Feet() -> Feet
func (m Meters) Feet() Feet {
	return Feet(m * FeetPerMeter)
}

// Meters.Yards() -> Yards
func (m Meters) Yards() Yards {
	return Yards(m * YardsPerMeter)
}

// DistFlags calls the flag package to setup distance's local vars.
func DistFlags() {
	ft = flag.Float64(ftFlag, math.Inf(1), "feet")
	m = flag.Float64(mFlag, math.Inf(1), "meters")
	y = flag.Float64(yFlag, math.Inf(1), "yards")
}

// DistConv checks the temp flags and returns a string with the requested
// conversion.
func DistConv(toScale *string) (string, error) {
	switch {
	case *ft != math.Inf(1) && *toScale == mFlag:
		// feet to meters
		return Feet(*ft).Meters().String(), nil
	case *y != math.Inf(1) && *toScale == mFlag:
		// yards to meters
		return Yards(*y).Meters().String(), nil
	case *m != math.Inf(1) && *toScale == ftFlag:
		// meters to feet
		return Meters(*m).Feet().String(), nil
	case *m != math.Inf(1) && *toScale == yFlag:
		// meters to yards
		return Meters(*m).Yards().String(), nil
	default:
		return "", errors.New("DistConv: unsupported temperature conversion")
	}
}
