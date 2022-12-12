package main

import (
	"car/parts"
	"fmt"
)

func main() {
	fmt.Println("Access a struct from an external package in Go.")

	// access the stuff using <packageName>.<Identifier>
	// newCar := parts.BluePrint{"Toyota", "someEngine", 4}
	newCar := parts.BluePrint{Name: "Toyota", Engine: "SomeEngine", Wheels: 4}

	fmt.Println(newCar)
}
