package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// creating a reader
	// bufio => read
	// os => where to read from
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// whatever the reader puts in, Trying to store that in a variable
	// Here is where COMMA OK / ERROR OK SYNTAX come into the picture
	// There are no try-catch in go
	// input, error => if care about error other wise put _ like shown below

	// \n => ender => keep on reading till there is a \n
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T \n", input)
}