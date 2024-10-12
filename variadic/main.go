package main

import "fmt"

func main() {
	max := maxInt(10000, 10, 100, 1000)

	fmt.Println(max)
}

func maxInt(nums ...int) int {
	var max int
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
