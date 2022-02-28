package reverse

import (
	"fmt"
	"reflect"
)

//!+rev
// reverse reverses a slice of ints in place.
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev



// reverse func using a array pointer as argument

func ReverseUsingPointer(s *[6]int){
	fmt.Println("refelct \n", reflect.ValueOf(s).Kind())
	str  := *s
	for i,j := 0, len(str)-1; i<j; i,j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	*s = str
}