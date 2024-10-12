package main

import "fmt"

type rectangle struct {
	w, l float64
}

func (rec rectangle) area() float64 {
	return rec.l * rec.w
}

func (rec *rectangle) setWidth(w float64) {
	rec.w = w
}

func main() {
	rec := rectangle{
		w: 4,
		l: 5,
	}

	rec.setWidth(10)
	fmt.Println("area = ", rec.area())
}
