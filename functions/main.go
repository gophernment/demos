package main

import "fmt"

func main() {
	fmt.Println(min(10, 2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
