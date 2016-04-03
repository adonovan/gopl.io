package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "Feburary", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 11: "November", 12: "December"}

	fmt.Println(months)

	// Print months by quarter
	//Q1 := months[1:4]
	//Q2 := months[4:7]
	//Q3 := months[7:10]
	//Q4 := months[10:13]

	fmt.Println(months[1:4])
	fmt.Println(months[4:7])
	fmt.Println(months[7:10])
	fmt.Println(months[10:13])

	// Print Months by Season
	Summer := months[6:9]
	// this isn't working
	// Trying to append 2 slices

	var Winter []string = string months[11:13] + string months[1:3]
	Spring := months[3:7]
	Fall := months[9:11]

	fmt.Println(Summer)
	fmt.Println(Fall)
	//This is not working. Trying to print [November, December, January, Feburary]
	// I get an error stating that I cannot add slices if I use a '+' operator
	//This was the closest I could get. It prints [November December] [January Feburary]
	fmt.Println(months[11:13], months[1:3])

	fmt.Println(Winter) // See above
	fmt.Println(Spring)

}
