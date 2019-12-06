package tempconv

import "fmt"

func Example_One() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
}

func Exmaple_two() {
	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c) // Printf automatically call String()
	fmt.Println(c)
	fmt.Printf("%g\n", c)   // does not call String()
	fmt.Println(float64(c)) // does not call String()

}
