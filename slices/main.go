package main

import "fmt"

func main() {
	var a = [...]int{-1, 0, 1, 2}
	var s []int

	s = a[1:3]

	fmt.Println(cap(s))
}
