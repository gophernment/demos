package main

import "fmt"

func main() {
	ch := make(chan int)
	qCH := make(chan struct{})

	go finacci(ch, qCH)

	for range 10 {
		fmt.Println(<-ch)
	}

	qCH <- struct{}{}

	<-qCH
}

func finacci(ch chan int, qCH chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-qCH:
			fmt.Println("bye")
			qCH <- struct{}{}
			return
		}
	}
}
