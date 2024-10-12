package main

import "fmt"

func main() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for i := range 5 {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

loopInfinity:
	for {
		if i < 10 {
			fmt.Println(i)
			i++
		} else {
			break loopInfinity
		}
	}
}
