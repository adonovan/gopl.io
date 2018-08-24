/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package conv

// weight conversions

import (
	"errors"
	"flag"
	"fmt"
	"math"
)

type Pounds float64
type Kilograms float64
type Stone float64

const (
	PoundsPerKilogram = 2.20462
	StonePerKilogram  = 0.157473
	kgFlag            = "kg"
	lbFlag            = "lb"
	stFlag            = "st"
)

var (
	kg, lb, st *float64
)

// Stringer interface
func (lb Pounds) String() string    { return fmt.Sprintf("%g pounds", lb) }
func (kg Kilograms) String() string { return fmt.Sprintf("%g kilograms", kg) }
func (st Stone) String() string     { return fmt.Sprintf("%g stone", st) }

// Pounds.Kilograms() -> Kilograms
func (lb Pounds) Kilograms() Kilograms {
	return Kilograms(lb / PoundsPerKilogram)
}

// Stone.Kilograms() -> Kilograms
func (st Stone) Kilograms() Kilograms {
	return Kilograms(st / StonePerKilogram)
}

// Kilograms.Pounds() -> Pounds
func (kg Kilograms) Pounds() Pounds {
	return Pounds(kg * PoundsPerKilogram)
}

// Kilograms.Stone() -> Stone
func (kg Kilograms) Stone() Stone {
	return Stone(kg * StonePerKilogram)
}

// WeightFlags calls the flag package to setup weight's local vars.
func WeightFlags() {
	kg = flag.Float64(kgFlag, math.Inf(1), "kilograms")
	lb = flag.Float64(lbFlag, math.Inf(1), "pounds")
	st = flag.Float64(stFlag, math.Inf(1), "stone")
}

// WeightConv checks the temp flags and returns a string with the requested
// conversion.
func WeightConv(toScale *string) (string, error) {
	switch {
	case *lb != math.Inf(1) && *toScale == kgFlag:
		// pounds to kilograms
		return Pounds(*lb).Kilograms().String(), nil
	case *st != math.Inf(1) && *toScale == kgFlag:
		// stone to kilograms
		return Stone(*st).Kilograms().String(), nil
	case *kg != math.Inf(1) && *toScale == lbFlag:
		// kilograms to pounds
		return Kilograms(*kg).Pounds().String(), nil
	case *kg != math.Inf(1) && *toScale == stFlag:
		// kilograms to stone
		return Kilograms(*kg).Stone().String(), nil
	default:
		return "", errors.New("WeightConv: unsupported temperature conversion")
	}
}
