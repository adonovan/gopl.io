package main

import "fmt"

func main()  {
	v := 0
	fmt.Println(incr(&v))
}

func incr(p *int) int  {
	w :=1
	p = &w
	fmt.Println("p ",p)
	*p++
	fmt.Println("*p ", *p)
	return *p
}