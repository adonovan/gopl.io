package tempconv

// CToF는 섭씨온도를 화씨온도로 변환한다.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC는 화씨온도를 섭씨온도로 변환한다.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
