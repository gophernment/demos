package main

import "fmt"

func main() {
	m := map[string]string{
		"d": "durian",
	}

	m["a"] = "apple"
	m["b"] = "banana"
	m["c"] = "coconut"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
