package tempconv

import "fmt"

// Celsius is float64
type Celsius float64

//Fahrenheit is float64
type Fahrenheit float64

const (
	// AbsoluteZeroC is coldest
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is Water Freezing
	FreezingC Celsius = 0
	// BoilingC is Water boiling
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
