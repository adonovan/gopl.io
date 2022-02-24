package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fmt.Println(6.0/3.5)
	fmt.Println(5/2)
	
	// read signed and unsigned int concept to undersatnd +1 and why result is printed as negetive.
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	f:= 3.14
	// f:= 1e100

	fmt.Printf("int = %d, float= %f \n",int(f), f)


	if i:=0; i == 0 {
		i++
		fmt.Println("i : ",i)
	}

	s := "Hello world"
	fmt.Println(s[0], s[7])
//	fmt.Printf("%c, %c",s[0], s[7])
	for i, v := range "Hello, ప్రపంచం" {
		fmt.Printf("%d \t %q \t %d\n", i, v, v)
	}

	j := "1"
	fmt.Println(strconv.Atoi(j))

}

func booleanToInteger(b bool) int {
	if b {
		return 1
	}
	return 0
}

func integerToBoolean(i int) bool{
	// if i !=0 {
	// 	return true
	// }
	// return false
	return i != 0
}