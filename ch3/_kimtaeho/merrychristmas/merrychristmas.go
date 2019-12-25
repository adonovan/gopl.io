package main

import "fmt"

func main() {

	for x := 0; x < 3; x++ {

		for i := 0; i < 4+x; i++ {
			for j := 0; j < 6-i; j++ {
				fmt.Printf(" ")
			}
			for j := 0; j < 2*i+1; j++ {
				fmt.Printf("*")
			}

			fmt.Println("")
		}

	}

	fmt.Println("Merry Christmas~*")

}
