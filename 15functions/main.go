package main

import "fmt"

func main() {
	fmt.Println("Welcome to the function tutorials in golang")
	greeter()

	// 
	// NOT ALLOWED TO WRITE FUNCTIONS INSIDE OF FUNCTION
	// 

	// func greeterTwo()  {
	// 	fmt.Println("Another method")
	// }

	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, _ := proAdder(2, 5, 8, 7)
	fmt.Println("proResult is: ", proResult)

	_, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println(myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// values now here is slice
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total,"Hi pro result function"
}

func greeterTwo() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Namaste from golang")
}