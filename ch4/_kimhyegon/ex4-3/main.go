package main

import "fmt"

func reverse(p *[]int) {
	s := p[0]
	//for i, j :=0, len(s)-1;  i<j; i,j = i+1, j-1 {
	//	s[i], s[j] = s[j], s[i]
	//}

	fmt.Printf("%v\n", s)
}

func main() {
	//a :=[...]int{0, 1, 2, 3, 4, 5}
	//reverse(&a[0])
	//fmt.Println(a)

	s := []int{1, 2, 3, 4, 5}
	sTemp := s[:2]
	reverse(&sTemp[0])
	sTemp = s[2:]
	reverse(&sTemp[0])
	reverse(&s[0])
	fmt.Println(&s[0])
}
