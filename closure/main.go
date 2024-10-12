package main

import "fmt"

func main() {
	isOdd := func(i int) bool {
		return i&1 == 1
	}
	isEven := func(i int) bool {
		return i&1 == 0
	}
	oddEven(1, isEven)
	oddEven(1, isOdd)
	oddEven(2, isEven)
	oddEven(2, isOdd)

	fn := newCounter()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

func oddEven(n int, fn func(int) bool) {
	fmt.Println(fn(n))
}

func newCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
