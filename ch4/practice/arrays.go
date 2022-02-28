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

	months := [...]string{
		1: "Jan",2: "Feb",3:"Mar",4:"April",5:"May",6:"June",7:"July",8:"Aug",9:"Sep",10:"Oct",11:"Nov",12:"Dec",
	}

	Q1 := months[1:4]
	Summer := months[3:7]
	fmt.Println(months[0], Q1, Summer)
	for _, summer := range Summer {
		for _, month := range Q1 {
			if month == summer {
				fmt.Printf("%s appears in both slices \n", month)
			}
		}
	}

	// fmt.Println(Summer[:20]) // throws error because, out of capacity of array months
	fmt.Println(Summer[0:7])
}