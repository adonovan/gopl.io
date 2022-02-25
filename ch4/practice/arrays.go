package main

import "fmt"


func main()  {
	//q := [3]int{1,2,3}
	//r := [4]int {1,2,3,4}
	const (
		USD  = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "#", GBP: "&", RMB: "@"}

	fmt.Println(symbol[1])
}