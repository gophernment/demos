package main

import "fmt"

func main() {
	concat := concatenate

	fmt.Println(concat("Hi,", "students"))

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	fmt.Println(max(0b1010001111, 0b0111101101))
}

func concatenate(a, b string) string {
	return a + b
}
