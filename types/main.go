package main

import "fmt"

func main() {
	var price = 9.5
	var paid = 10

	change := float64(paid) - price
	fmt.Println(change)
}
