package main

import "fmt"

func main() {
	var name []string
	type Car struct {
		name  string
	}
	
	type CarTwo struct {
	}



	myCar := Car{
		name:  "bmv",
		model: 1244,
	}
	fmt.Println(myCar.model,name)
}
