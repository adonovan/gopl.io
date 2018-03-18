/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package conv

// temperature conversions

import (
	"errors"
	"flag"
	"fmt"
	"math"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15

	FreezingC Celsius = 0
	BoilingC  Celsius = 100

	FreezingK Kelvin = 273.15
	BoilingK  Kelvin = 373.15

	FreezingF Fahrenheit = 32
	BoilingF  Fahrenheit = 212

	fFlag = "f"
	cFlag = "c"
	kFlag = "k"
)

var (
	f, c, k *float64
)

// Stringer interface
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

// Celsius.Fahrenheit() -> Fahrenheit
func (c Celsius) Fahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Fahrenheit.Celsius() -> Celsius
func (f Fahrenheit) Celsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Kelvin.Celsius() -> Celsius
func (k Kelvin) Celsius() Celsius {
	return Celsius(k) + Celsius(float64(AbsoluteZeroC))
}

// Celsius.Kelvin() -> Kelvin
func (c Celsius) Kelvin() Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

// TempFlags calls the flag package to setup temperature's local vars.
func TempFlags() {
	f = flag.Float64(fFlag, math.Inf(1), "fahrenheight")
	c = flag.Float64(cFlag, math.Inf(1), "celsius")
	k = flag.Float64(kFlag, math.Inf(1), "kelvin")
}

// TempConv checks the temp flags and returns a string with the requested conversion.
func TempConv(toScale *string) (string, error) {
	switch {
	case *f != math.Inf(1) && *toScale == cFlag:
		// fahrenheit to celsius
		return Fahrenheit(*f).Celsius().String(), nil
	case *c != math.Inf(1) && *toScale == fFlag:
		// celsius to fahrenheit
		return Celsius(*c).Fahrenheit().String(), nil
	case *c != math.Inf(1) && *toScale == kFlag:
		// celsius to kelvin
		return Celsius(*c).Kelvin().String(), nil
	case *k != math.Inf(1) && *toScale == cFlag:
		// kelvin to celsius
		return Kelvin(*k).Celsius().String(), nil
	default:
		return "", errors.New("TempConv: unsupported temperature conversion")
	}
}
