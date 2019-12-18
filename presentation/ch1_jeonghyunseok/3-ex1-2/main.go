// 1.2 echo 프로그램을 수정해 각 인자의 인덱스와 값을 한 줄에 하나씩 출력하라
// start from echo3 example
// ex1-2 will print index and value of the arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("Index\tArgument")
	fmt.Println("------------------")
	for i, v := range os.Args[1:] {
		fmt.Printf("%d\t%s\n", i+1, v)
	}
}

/*
go build -o ex1-2.exe
.\ex1-2.exe Who let the dogs out!
*/
