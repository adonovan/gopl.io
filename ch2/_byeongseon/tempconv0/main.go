// tempconv 패키지는 섭씨와 화씨온도는 연산을 수행한다.
package main

import "fmt"

// Celsius is float64
type Celsius float64

// Fahrenheit is float64
type Fahrenheit float64

const (
	// AbsoluteZeroC is coldest
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is Water Freezing
	FreezingC Celsius = 0
	// BoilingC is Water boiling
	BoilingC Celsius = 100
)

func main() {
	fmt.Printf("%g°C\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g°F\n", boilingF-CToF(FreezingC))
	// 컴파일 오류 : 타입 불일치
	// fmt.Printf("%g\n", boilingF, FreezingC)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// 컴파일 오류 : 타입 불일치
	// fmt.Println(c == f)
	fmt.Println(c == Celsius(f))

	ce := FToC(212.0)
	fmt.Println(ce.String())

	// String을 명시적으로 호출할 필요가 없음
	// 왜일까...?
	fmt.Printf("%v\n", ce)
	fmt.Printf("%s\n", ce)
	fmt.Println(ce)

	// String을 호출하지 않음
	fmt.Printf("%g\n", ce)
	fmt.Println(float64(ce))
}

// CToF is CToF
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC is FToC
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
