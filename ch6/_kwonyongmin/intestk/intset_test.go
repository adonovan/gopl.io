package intestk

import "fmt"

func Example_one() {
	var x, y intest
	x.Add(1)
	x.Add(144)
	x.Add(9)

	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)

	fmt.Println(y.String())

	x.UnionWith(&y)

	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}

func Example_two() {
	var x intest

	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)
	fmt.Println(x.String())
	fmt.Println(x)
}
