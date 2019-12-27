package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = 273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	//fmt.Print("%g\n", boilingF - FreezingC) // 컴파일 오류 : 타입 불일치

	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) // 컴파일 오류: 타입 불일치
	fmt.Println(c == Celsius(f))

	{
		c := FToC(212.0)
		fmt.Println(c.String())
		fmt.Println("%v\n", c)
		fmt.Println("%s\n", c)
		fmt.Println(c)
		fmt.Println("%g\n", c)
		fmt.Println(float64(c))
	}
}
