package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) { // 확장할 공간이 있음
		z = x[:zlen]
	} else { // 두배로 증가시킨 새 배열 할당
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d	cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

/* output
0	cap=1	[0]
1	cap=2	[0 1]
2	cap=4	[0 1 2]
3	cap=4	[0 1 2 3]
4	cap=8	[0 1 2 3 4]
5	cap=8	[0 1 2 3 4 5]
6	cap=8	[0 1 2 3 4 5 6]
7	cap=8	[0 1 2 3 4 5 6 7]
8	cap=16	[0 1 2 3 4 5 6 7 8]
9	cap=16	[0 1 2 3 4 5 6 7 8 9]
*/
