package main

import (
	"fmt"
	"sync"
	"time"
)

type dataMux struct {
	i   int
	mux sync.RWMutex
	wg  sync.WaitGroup // WaitGroup ใช้เพื่อ sync ให้ผลลัพธ์กลับมาครบ
}

var data dataMux

func main() {

	start := time.Now()
	for range 5 {
		data.wg.Add(1)
		go slow(func() {
			fmt.Println(data.i)
		})
	}

	data.wg.Wait()

	fmt.Println(time.Since(start))
}

func slow(fn func()) {
	time.Sleep(time.Millisecond * 100)
	data.mux.Lock()
	data.i++
	data.mux.Unlock()
	fn()
	data.wg.Done()
}
