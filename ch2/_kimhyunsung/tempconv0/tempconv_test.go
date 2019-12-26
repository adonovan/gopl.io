package tempconv

import "fmt"

func Example_one() {
	{
		fmt.Printf("%g\n", BoilingC-FreezingC)
		boilingF := CToF(BoilingC)
		fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	}
}

func Example_two() {
	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
