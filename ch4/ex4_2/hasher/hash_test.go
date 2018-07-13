/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package hasher

import (
	"testing"
)

// TestHash* tests some results generated by osx's shasum cli tool
func TestHash256(t *testing.T) {
	got := Hash(Sha256, []byte("alex\n"))
	want := "93191218324b205ca4e4c0e0a9600759c818f9b908ec9311e9147178701ba04d"
	if got != want {
		t.Errorf("Hash sha256 = %s, want %s", got, want)
	}
}

// TestHash* tests some results generated by osx's shasum cli tool
func TestHash384(t *testing.T) {
	got := Hash(Sha384, []byte("alex\n"))
	want := "8dbfc82a64a953881a06177235caf4340066b16b44e448bec0d0175c348fb08dd06ac2d6baf58402adeae5f5221430fc"
	if got != want {
		t.Errorf("Hash Sha384 = %s, want %s", got, want)
	}
}

// TestHash* tests some results generated by osx's shasum cli tool
func TestHash512(t *testing.T) {
	got := Hash(Sha512, []byte("alex\n"))
	want := "3570922399a3a278330222a218aadb4de26fd504eedfb38512e408182d2bfd490cb8d0df92b32c0ab04ee3740257ed8a50843a1151e7cd1d9bee8e002d76cc21"
	if got != want {
		t.Errorf("Hash Sha512 = %s, want %s", got, want)
	}
}
