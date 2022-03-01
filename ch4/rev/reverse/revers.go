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


func Rotate(arr []int, n int, direction bool)  {
	fmt.Printf("arr : %v, n: %d \n", arr, n)
	// direction = true => rotate to right ,
	// direction = false => rotate to left
	if direction {
		tempArr := arr
		arr = tempArr[len(tempArr)-n:]
		arr = append(arr, tempArr[0:(len(tempArr)-n)]...)
		fmt.Println(arr)
	}else{
		tempArr := arr
		arr = tempArr[n:]
		arr = append(arr, tempArr[0:n]...)
		fmt.Println(arr)
	}
}

func RemoveDuplicatesFromStringSlice(arr []string)  {
	for i, a := range arr {
		if i < len(arr)-1 && a == arr[i+1]{
			arr[i] = ""
		}
		fmt.Println(arr)
	}
	fmt.Println("RemoveDuplicatesFromStringSlice: ", nonempty(arr))
}

func nonempty(strings []string) []string {
	fmt.Println(strings)
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		fmt.Printf("s: %s, string: %s\n",s, strings)

		}else{
		//fmt.Println("in else ",s)
		}
	}
	//fmt.Printf("string: %s\n",strings)
	return strings[:i]
}


func ReverseTheCharOfByteSlice(b []byte)  {
	for i,j := 0, len(b)-1; i < j; i,j =i+1,j-1 {
		b[i],b[j] = b[j], b[i]
	}
	fmt.Println(string(b))
}