package main

import "fmt"

type number int

func (n number) String() string {
	return fmt.Sprintf("%d", int(n))
}

type rectangle struct {
	w, l float64
}

func (rec rectangle) area() float64 {
	return rec.l * rec.w
}

func main() {
	var year number = 2024
	s := "this year is " + year.String()
	fmt.Println(s)

	rec := rectangle{
		w: 4,
		l: 5,
	}

	fmt.Println("area = ", rec.area())
}
