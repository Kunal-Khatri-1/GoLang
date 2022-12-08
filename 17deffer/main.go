package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("World 2")
	defer fmt.Println("World 1")
	fmt.Println("End")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}