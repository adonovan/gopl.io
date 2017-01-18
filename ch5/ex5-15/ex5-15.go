// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.

//ex5-15
//added a Max and Min functions to find the max and min values of the data set
//in min, I first used the max function to find the max value and then
//used that result as the starting point to find min.
//Both will print a message if no values are provided inthe data set.

package main

import "fmt"

//!+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func maximum(vals ...int) int {

	if len(vals) == 0 {
		fmt.Println("Must supply at least one value!")
		return 0
	}
	max := 0
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func minimum(vals ...int) int {

	if len(vals) == 0 {
		fmt.Println("Must supply at least one value!")
		return 0
	}

	min := maximum(vals...)

	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

//!-

func main() {
	//!+main
	fmt.Println("Sum:")
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	//!-slice

	//!+main
	fmt.Println("Max:")
	fmt.Println(maximum())            //  "0"
	fmt.Println(maximum(3))           //  "3"
	fmt.Println(maximum(1, 20, 3, 4)) //  "10"
	//!-main

	//!+slice
	values = []int{1, 2, 30, 4}
	fmt.Println(maximum(values...)) // "10"
	//!-slice

	//!+main
	fmt.Println("Min:")
	fmt.Println(minimum())              //  "0"
	fmt.Println(minimum(3))             //  "3"
	fmt.Println(minimum(1, 20, 30, 40)) //  "10"
	//!-main

	//!+slice
	values = []int{10, 20, 30, 40}
	fmt.Println(minimum(values...)) // "10"
	//!-slice
}
