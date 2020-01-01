package main_test

import (
	"fmt"
	"strings"
	"testing"
)

var args = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

func echo1() {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	fmt.Println(s)
}

func echo2() {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(args[0:], " "))
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}
