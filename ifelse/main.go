package main

import (
	"fmt"
	"log"
	"time"
)

/*
Generations	Born	Current Ages
Gen Z	1997 – 2012	12 – 27
Millennials	1981 – 1996	28 – 43
Gen X	1965 – 1980	44 – 59
Boomers II (a/k/a Generation Jones)*	1955 – 1964	60 – 69
Boomers I*	1946 – 1954	70 – 78
Post War	1928 – 1945	79 – 96
WWII	1922 – 1927	97 – 102

yearOfBirth > 2012: Young
yearOfBirth >= 1997: Gen Z
yearOfBirth >=1981: Millennials
else: Senior
*/

func main() {
	var age int
	_, err := fmt.Scanln(&age)
	if err != nil {
		log.Panic(err)
	}

	thisYear := time.Now().Year()
	yearOfBirth := thisYear - age

	if yearOfBirth > 2012 {
		fmt.Println("you are too young")
		return
	}

	if yearOfBirth >= 1997 {
		fmt.Println("Gen Z")
		return
	}

	if yearOfBirth >= 1981 {
		fmt.Println("Millennials")
		return
	}

	if yearOfBirth >= 1965 {
		fmt.Println("Gen X")
	} else {
		fmt.Println("Hi senior")
	}
}

func switcaseVersion() {
	var age int
	_, err := fmt.Scanln(&age)
	if err != nil {
		log.Panic(err)
	}

	year := time.Now().Year()
	yearOfBirth := year - age

	switch {
	case yearOfBirth > 2012:
		fmt.Println("young")
	case yearOfBirth >= 1997:
		fmt.Println("Gen Z")
	default:
		fmt.Println("Hi Senior")
	}

}
