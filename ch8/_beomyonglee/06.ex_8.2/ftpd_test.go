package main

import (
	"testing"
)

func TestHostPortToFTP(t *testing.T) {
	addr, err := hostPortToFTP("192.168.1.34:3243")
	if err != nil {
		t.Fatal(err)
	}
	want := "192,168,1,34,12,171"
	if addr != want {
		t.Fatalf("got=%s want=%s", addr, want)
	}
}

func TestHostPortFromFTP(t *testing.T) {
	addr, err := hostPortFromFTP("192,168,1,34,12,171")
	if err != nil {
		t.Fatal(err)
	}
	want := "192.168.1.34:3243"
	if addr != want {
		t.Fatalf("got=%s want=%s", addr, want)
	}
}
