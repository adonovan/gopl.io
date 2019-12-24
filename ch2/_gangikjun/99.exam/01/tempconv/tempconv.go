// Package tempconv 는 온도의 변환에 관련된 것들을 정리
package tempconv

import "fmt"

// Celsius 섭씨
type Celsius float64

// Fahrenheit 화씨
type Fahrenheit float64

// Kelvin 켈빈
type Kelvin float64

// 온도에 관련된 상수
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
