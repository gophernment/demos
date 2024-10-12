package main

import (
	"fmt"
	"time"
)

func main() {
	a := ransomThing()

	decision(a)
}

func decision(a any) {
	switch t := a.(type) {
	case string:
		fmt.Printf("%T: %v\n", t, a)
	case int:
		fmt.Printf("%T: %v\n", t, a)
	case float64:
		fmt.Printf("%T: %v\n", t, a)
	default:
		fmt.Printf("%T %v\n", t, a)
	}
}

func ransomThing() any {
	n := time.Now().Second()

	switch n % 10 {
	case 0:
		return &n
	case 1:
		return 10
	case 2:
		return "ten"
	case 3:
		return 10.001
	case 4:
		return true
	case 5:
		return uint8(10)
	case 6:
		return func() int { return 10 }
	case 7:
		return map[string]int{"ten": 10}
	case 8:
		return []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	case 9:
		return struct{}{}
	}
	return int32(0)
}
