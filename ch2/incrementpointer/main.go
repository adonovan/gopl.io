package main

import "fmt"

func main()  {
	v := 0
	incr(&v)
	fmt.Println(v)
	incrWithNewVariable(&v)
	fibonacci(6)

	// below code is to verify if the variable has different 
	// address if declared in different blocks.
	// ans is yes they do have different address.
	x := "h"
	fmt.Println(&x)
	for _, x := range x {
		fmt.Println(&x)
		fmt.Printf( "%c\n",x)
	}
}

func incr(p *int) int  {
	*p++
	return *p
}


func incrWithNewVariable(p *int) int  {
	w :=1
	p = &w
	//fmt.Println("p ",p)
	*p++
	fmt.Println("*p ", *p)
	return *p
}


//nth fibonacci number
func fibonacci(n int) int {
	x,y := 0,1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	fmt.Println("nth fibonacci number is ",x)
	return x
}