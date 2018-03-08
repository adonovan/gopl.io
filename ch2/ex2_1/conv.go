// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

// ex 2.1 add types, constants, and functions to tempconv for processing
// temperatures in Kelvin scale, where zero kelvin is -273.13 °C and a
// difference of 1K has the same magnitude as 1°C.

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func KToC(k Kelvin) Celsius {
	return Celsius(k) + Celsius(float64(AbsoluteZeroC))
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}
