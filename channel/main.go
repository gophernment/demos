package main

import (
	"fmt"
)

func main() {
	demoNoBufferChannel()
}

// การทำให้ unbuffered channel ทำงานได้ปกติ จะต้องมี 2 goroutine ทำงานพร้อมกัน
// ทั้งฝั่งส่งและฝั่งรับ
func demoNoBufferChannel() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)
}

// เมื่อใช้ unbuffered channel การจะส่งค่าเข้าไป จำเป็นจะต้องมีอีก goroutine หนึ่งมารับค่าออกไปด้วย
// มิเช้านั้น ก็จะติด deadlock
func demoNoBufferDeadlock() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}

// channel ที่ถูก close หมายถึงจะไม่สามารถส่งค่าใดเข้าไปได้อีก
// ส่งผลให้เมื่อต้องการดึงค่าออกมา จะสามารถดึงได้ แต่จะได้ค่า zero value ของ channel ออกมา
func demoClose() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// เมื่อต้องการส่งค่าเข้าไปใน channel จะต้องทำให้ channel มี buffer เหลือเพียงพอก่อน ด้วยการนำค่าออกไป
func demoUnlock() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	ch <- 3

	fmt.Println(<-ch)
}

// deadlock เพราะมี 2 buffer แต่พยายามจะดึงค่าที่ 3 โดยที่ไม่มีค่าที่ 3 ถูกส่งเข้าไปใน channel
func demoDeadlockReceive() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// deadlock เพราะมี 2 buffer แต่พยายามจะส่งของเข้าไป 3 ค่า ซึ่งเกินปริมาณ buffer
func demoDeadlockSend() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// bufferd channel ตัวอย่างมี buffer 2 int
func demoBufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// ตัวอย่างการเรียก fibonacci
func callFibonacci() {
	ch := make(chan int)
	go fibonacci(ch)

	for range 10 {
		fmt.Print(<-ch, " ")
	}
}

// ตัวอย่างของการสร้าง fibonacci sequence และส่งค่ากลับไปทาง channel
func fibonacci(ch chan int) {
	a, b := 0, 1
	for {
		ch <- a
		a, b = b, a+b
	}
}
