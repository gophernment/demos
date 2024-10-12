package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	slow()
}

func fastSync() {
	wg := sync.WaitGroup{}
	start := time.Now()
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

func fastNotSync() {
	start := time.Now()
	for i := range 5 {
		go delay(i)
	}
	fmt.Println(time.Since(start))
}

func slow() {
	start := time.Now()
	for i := range 5 {
		delay(i)
	}
	fmt.Println(time.Since(start))
}

func delay(i int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(i)
}
