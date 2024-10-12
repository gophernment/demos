package main

import (
	"fmt"
)

type ascii uint8
type char = uint8

func main() {
	var c char = 1
	var a ascii = 1

	fmt.Println(ascii(c) + a)
}
