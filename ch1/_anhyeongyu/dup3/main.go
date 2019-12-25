package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// data -> []byte
		// 파일의 모든 데이터를 가져오기 때문에
		// 많은 메모리를 차지할 위험성이 있음!!
		data, err := ioutil.ReadFile(filename)
		// nil은 키워드 인가?
		// No!!
		// Reference type value
		// 키워드란
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
