/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package conv

import (
	"log"
	"testing"
)

func TestLbToKg(t *testing.T) {
	want := Kilograms(1)
	got := Pounds(PoundsPerKilogram).Kilograms()
	if got != want {
		log.Fatalf("Kilograms: got %v, want %v", got, want)
	}
}

func TestStToKg(t *testing.T) {
	want := Kilograms(1)
	got := Stone(StonePerKilogram).Kilograms()
	if got != want {
		log.Fatalf("Kilograms: got %v, want %v", got, want)
	}
}

func TestKgToLb(t *testing.T) {
	want := Pounds(PoundsPerKilogram)
	got := Kilograms(1).Pounds()
	if got != want {
		log.Fatalf("Pounds: got %v, want %v", got, want)
	}
}

func TestKgToSt(t *testing.T) {
	want := Stone(StonePerKilogram)
	got := Kilograms(1).Stone()
	if got != want {
		log.Fatalf("Stone: got %v, want %v", got, want)
	}
}
