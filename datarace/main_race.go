package main

import (
	"fmt"
	"time"
)

var i int

func main() {

	start := time.Now()
	for range 5 {
		go slow(func() {
			fmt.Println(i)
		})
	}

	for {
		if i >= 4 {
			break
		}
	}

	fmt.Println(time.Since(start))
}

func slow(fn func()) {
	time.Sleep(time.Millisecond * 100)
	i++
	fn()
}
