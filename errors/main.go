package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Name string
}

func main() {
	a := Account{
		Name: "Skooldio",
	}

	b, err := json.Marshal(&a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(b[0])

}
