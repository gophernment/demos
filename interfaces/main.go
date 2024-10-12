package main

import "fmt"

type dog string

func (d dog) Sound() string {
	return "โฮ่ง"
}

type cat string

func (c cat) Sound() string {
	return "เมี๊ยว"
}

type animal interface {
	Sound() string
}

func playSound(a animal) {
	fmt.Println(a.Sound())
}

func main() {
	var d dog
	playSound(d)

	var c cat
	playSound(c)
}
