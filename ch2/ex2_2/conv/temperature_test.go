/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package conv

import (
	"log"
	"testing"
)

func TestCToF(t *testing.T) {
	want := FreezingF
	got := FreezingC.Fahrenheit()
	if got != want {
		log.Fatalf("Fahrenheit: got %v, want %v", got, want)
	}
}

func TestFToC(t *testing.T) {
	want := FreezingC
	got := FreezingF.Celsius()
	if got != want {
		log.Fatalf("Celsius: got %v, want %v", got, want)
	}
}

func TestKToC(t *testing.T) {
	want := AbsoluteZeroC
	got := Kelvin(0).Celsius()
	if got != want {
		log.Fatalf("Celsius: got %v, want %v", got, want)
	}
}

func TestCToK(t *testing.T) {
	want := Kelvin(0)
	got := AbsoluteZeroC.Kelvin()
	if got != want {
		log.Fatalf("Kelvin: got %v, want %v", got, want)
	}
}

func TestBoiling(t *testing.T) {
	want := BoilingK
	got := want.Celsius()
	if got != BoilingC {
		log.Fatalf("Celsius: got %v, want %v", got, want)
	}
	f := got.Fahrenheit()
	if f != BoilingF {
		log.Fatalf("Fahrenheit: got %v, want %v", got, want)
	}
}

func TestFreezing(t *testing.T) {
	want := FreezingK
	got := want.Celsius()
	if got != FreezingC {
		log.Fatalf("Celsius: got %v, want %v", got, want)
	}
	wantF := FreezingF
	gotF := FreezingC.Fahrenheit()
	if gotF != wantF {
		log.Fatalf("Fahrenheit: got %v, want %v", wantF, wantF)
	}
}
