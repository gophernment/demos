package main

import "fmt"

func main() {
	var a any

	a = 10
	fmt.Printf("type: %T, value: %v\n", a, a)
	a = "ten"
	fmt.Printf("type: %T, value: %v\n", a, a)
	a = true
	fmt.Printf("type: %T, value: %v\n", a, a)
	a = func() string { return "hello" }
	fmt.Printf("type: %T, value: %v\n", a, a)
}
