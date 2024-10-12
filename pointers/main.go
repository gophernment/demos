package main

import "fmt"

func main() {
	var p *int
	var n int = 14

	p = &n

	*p++

	fmt.Printf(" p: [%14p], &p: [%p]\n", p, &p)
	power(p)

	fmt.Println(n)
}

func power(p *int) {
	fmt.Printf(" p: [%14p], &p: [%p]\n", p, &p)
	*p = *p * *p
}
