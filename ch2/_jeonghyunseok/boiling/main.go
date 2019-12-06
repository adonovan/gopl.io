// Boiling is
package main

import "fmt"

const boilingF = 212.0 // 끓는 점 화씨

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gºF, or %gºC\n", f, c) // %g is %f or %e
	// boiling point = 212°F or 100°C
}

// go build -o boiling.exe
// boiling.exe
