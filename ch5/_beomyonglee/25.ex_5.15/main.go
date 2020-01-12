package main

import "fmt"

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func min2(first int, nums ...int) int {
	m := first
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func max2(first int, nums ...int) int {
	m := first
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func main() {
	fmt.Println(min(3, -1, 4))
	fmt.Println(max(3, -1, 4))
	fmt.Println(min2(3, -1, 4))
	fmt.Println(max2(3, -1, 4))
}
