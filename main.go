package main

import "fmt"

func main() {

	type Car struct {
		name  string
		model int
	}

	myCar := Car{
		name:  "bmv",
		model: 1244,
	}
fmt.Println(myCar)
}
