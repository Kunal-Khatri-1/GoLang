package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// * => this data type is of pointer, a pointer which going to be responsibe for holding integers
	// if you start a pointer don't intiialize it, then it will hold nil
	// var ptr *int	// ptr is nil here

	// fmt.Println("Value of ptr is: ", ptr)

	myNumber := 23
	// creating a poninter that points to myNumber memory address
	// & => used for referencing
	var ptr2 = &myNumber

	// pointer is ther reference to the direct memory location
	fmt.Println("Value of actual pointer is: ", ptr2)	// gives the memory address 0xc00...98
	// Value insid the pointer is 23
	// *ptr2 => not memory address but the value
	fmt.Println("Value of actual pointer is: ", *ptr2)	// get the value 23

	*ptr2 = *ptr2 * 2
	fmt.Println("New value is: ", myNumber)
}