package tempconv

import (
	"log"
	"testing"
)

func TestCToF(t *testing.T) {
	want := FreezingF
	f := CToF(FreezingC)
	if f != want {
		log.Fatalf("CToF: got %v, want %v", f, want)
	}
}

func TestFToC(t *testing.T) {
	want := FreezingC
	f := FToC(FreezingF)
	if f != want {
		log.Fatalf("FToC: got %v, want %v", f, want)
	}
}

func TestKToC(t *testing.T) {
	want := AbsoluteZeroC
	k := KToC(0)
	if k != want {
		log.Fatalf("KToC: got %v, want %v", k, want)
	}
}

func TestCToK(t *testing.T) {
	want := Kelvin(0)
	k := CToK(AbsoluteZeroC)
	if k != want {
		log.Fatalf("CToK: got %v, want %v", k, want)
	}
}

func TestBoiling(t *testing.T) {
	k := BoilingK
	c := KToC(k)
	if c != BoilingC {
		log.Fatalf("KToC: got %v, want %v", k, BoilingC)
	}
	f := CToF(c)
	if f != BoilingF {
		log.Fatalf("CToF: got %v, want %v", k, BoilingF)
	}
}

func TestFreezing(t *testing.T) {
	k := FreezingK
	c := KToC(k)
	if c != FreezingC {
		log.Fatalf("KToC: got %v, want %v", k, FreezingC)
	}
	f := CToF(c)
	if f != FreezingF {
		log.Fatalf("CToF: got %v, want %v", k, FreezingF)
	}
}
