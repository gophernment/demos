package main

import "fmt"

type rect struct {
	w, d float64
}

func (r rect) area() float64 {
	return r.d * r.w
}

func (r *rect) setWidth(w float64) {
	r.w = w
}

func main() {
	r := rect{
		w: 3,
		d: 4,
	}

	r.setWidth(5)

	fmt.Println(r.area())
}
