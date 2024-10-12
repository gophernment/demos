package main

import "fmt"

func main() {
	deferRecovery()
}

// ตัวอย่างนี้แสดงให้เห็นลำดับการทำงานของ defer
func deferConstant() {
	defer fmt.Println(1)
	fmt.Println(2)
}

// ตัวอย่างนี้แสดงให้เห็นว่าถ้ามีการ defer และมีการส่งค่า variable เข้าไปด้วย จะมีผลอย่างไร
func deferVariable() {
	i := 1
	defer fmt.Println(i)
	i++
	fmt.Println(i)
}

// ตัวอย่างนี้แสดงให้เห็นว่า ถ้า defer ด้วย anonymous function โดยไม่ได้ส่ง parameter ลงไป function ที่ถูก defer
// การ evaluate ค่า variable จะเกิดเมื่อ anonymous ถูกเรียกใช้งานจริง ทำให้ค่า i จะเป็นค่าที่เปลี่ยนไปแล้ว
func deferAnonymousFunction() {
	i := 1
	defer func() {
		fmt.Println(i)
	}()
	i++
	fmt.Println(i)
}

// ตัวอย่างการใช้ defer ทำ recovery
func deferRecovery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovery: %s\n", r)
		}
	}()
	fmt.Println([]int{}[0])
}
