package main

import (
	"fmt"
	"sync"
	"time"
)

// version นี้ ถึงจะแก้เรื่อง data race แล้ว แต่ก็ยังไม่รับประกันว่าจะได้ผลลัพธ์กลับมาครบ
// เนื่องจากใช้วิธีวนลูปนับค่า i ทำให้มีโอกาสที่ค่า i จะมีค่าเป็น 4 แต่อาจจะยังไม่ได้ print ค่า i ออกมา

type dataMux struct {
	i   int
	mux sync.RWMutex
}

var data dataMux

func main() {

	start := time.Now()
	for range 5 {
		go slow(func() {
			fmt.Println(data.i)
		})
	}

	for {
		data.mux.RLock()
		if data.i >= 4 {
			break
		}
		data.mux.RUnlock()
	}

	fmt.Println(time.Since(start))
}

func slow(fn func()) {
	time.Sleep(time.Millisecond * 100)
	data.mux.Lock()
	data.i++
	data.mux.Unlock()
	fn()
}
