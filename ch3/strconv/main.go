package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println(strconv.FormatInt(int64(x), 8))
	fmt.Println(strconv.FormatInt(int64(x), 10))
	fmt.Println(strconv.FormatInt(int64(x), 16))

	s := fmt.Sprintf("Binary x=%b\n", x)
	fmt.Println(s)

	s = fmt.Sprintf("Decimal x=%d\n", x)
	fmt.Println(s)

	s = fmt.Sprintf("U x=%u\n", x)
	fmt.Println(s)

	s = fmt.Sprintf("Hex x=%x\n", x)
	fmt.Println(s)

	x, err := strconv.Atoi("123d")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Atoi x=%d", x)
	}

}
