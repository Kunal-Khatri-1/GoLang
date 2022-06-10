package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:") // println has \n at the end by default

	// comma ok || err ok

	// _ => we don't care about this variable
	// if everything goes right then put stuff in input
	// we can use _ in place of input if we don't care about it
	input, _ := reader.ReadString('\n') // read as long as there is a \n
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T ", input)
}
