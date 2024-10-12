package main

import "fmt"

const (
	byte = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
	eb
	zb
	yb
	// Add more constants as needed...
)

func main() {
	fmt.Println(byte)
	fmt.Println(kb)
}
