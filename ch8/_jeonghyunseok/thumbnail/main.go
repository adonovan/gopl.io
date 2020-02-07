// 썸네일을 만들어준다.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopl.io/ch8/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text()) // 이게 실제 동작시키는 녀석
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
		// test
	}
}

/*
go build -o t.exe
t foo.jpeg


*/
