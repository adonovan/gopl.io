package tempconv

// Celsius 섭씨
type Celsius float64

// Fahrenheit 화씨
type Fahrenheit float64

// 온도에 관련된 상수
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF 섭씨를 화씨로
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC 화씨를 섭씨로
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
