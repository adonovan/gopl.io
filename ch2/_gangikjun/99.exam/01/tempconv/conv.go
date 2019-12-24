package tempconv

// CToF 섭씨를 화씨로
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC 화씨를 섭씨로
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC 켈빈을 섭씨로
func KToC(k Kelvin) Celsius {
	return Celsius(AbsoluteZeroC + Celsius(k))
}

// CToK 섭씨를 켈빈으로
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}
