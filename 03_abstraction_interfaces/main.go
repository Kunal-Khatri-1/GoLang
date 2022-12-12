package main

import "fmt"

type bird interface {
	fly()
}

type eagle struct {
	name string
}

type sparrow struct {
	name string
}

func (e eagle) fly() {
	fmt.Printf("Eagles can fly")
}

func (s sparrow) fly() {
	fmt.Println("Sparrows can fly")
}

func main() {
	fmt.Println("Abstraction using Interfaces in Golang")
	fmt.Println("Creating inerface that enables common abstraction that can be used by multiple types")

	eagleBird := eagle{name: "Eagle"}
	sparrowBird := sparrow{name: "Sparrow"}

	// creating a variable bluePrint of type bird Interface
	var bluePrint1 bird = eagleBird
	var bluePrint2 bird = sparrowBird
	fmt.Println("eagleBird", eagleBird)
	fmt.Println("bluePrint1", bluePrint1)

	bluePrint2.fly()
	bluePrint1.fly()
}
