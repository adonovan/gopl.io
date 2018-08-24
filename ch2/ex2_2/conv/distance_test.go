/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package conv

import (
	"log"
	"testing"
)

func TestFtToM(t *testing.T) {
	want := Meters(1)
	got := Feet(FeetPerMeter).Meters()
	if got != want {
		log.Fatalf("Meters: got %v, want %v", got, want)
	}
}

func TestYToM(t *testing.T) {
	want := Meters(1)
	got := Yards(YardsPerMeter).Meters()
	if got != want {
		log.Fatalf("Meters: got %v, want %v", got, want)
	}
}

func TestMtoF(t *testing.T) {
	want := Feet(FeetPerMeter)
	got := Meters(1).Feet()
	if got != want {
		log.Fatalf("Feet: got %v, want %v", got, want)
	}
}

func TestMtoY(t *testing.T) {
	want := Yards(YardsPerMeter)
	got := Meters(1).Yards()
	if got != want {
		log.Fatalf("Yards: got %v, want %v", got, want)
	}
}
