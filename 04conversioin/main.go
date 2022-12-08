package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// 
	// adding 1 to whatever rating user is giving
	// 

	// numRating = input +1		// this gives error

	// this takes 4\n and not 4 as input
	// numRating, err := strconv.ParseFloat(input, 64)

	// trimming whitespaces
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// everytime there is an err
	// move the program cautiously checking if there is any err or not

	// err is not empty => there is something in the err => there is an err
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}