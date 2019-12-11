package cc

import (
	"fmt"
	"strings"
)

// func main() {
// 	Concat(os.Args)
// 	Join(os.Args)
// }

func Join(src []string) {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	s := strings.Join(src, " ")
	fmt.Println(s)
}

func Concat(src []string) {
	var s, sep string
	for i := 1; i < len(src); i++ {
		s += sep + src[i]
		sep = " "
	}
	fmt.Println(s)
}
