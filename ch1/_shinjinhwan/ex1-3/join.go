package echo13

import (
	"fmt"
	"strings"
)
const text = "I am not your sounds hard it goes not out sound"

var textArray = strings.Split(text, " ")

func joinmethod(text []string) {
	fmt.Println(strings.Join(text,  textArray))
}

