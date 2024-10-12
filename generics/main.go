package main

import "fmt"

type number interface {
	~int | float64
}

func max[T number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type numeric int

func main() {
	var a, b float64 = 10, 99

	var maximum = max[float64]

	fmt.Println(maximum(a, b))

	fn := func() int {
		return 10
	}

	fmt.Println(fn())
}
