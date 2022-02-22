package main

import "fmt"

func main()  {
	v := 0
	incr(&v)
	fmt.Println(v)
}

func incr(p *int) int  {
	*p++
	return *p
}