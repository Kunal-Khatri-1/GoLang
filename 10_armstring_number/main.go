package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Program to find Armstrong Number")

	// creating reader to get input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Number: ")
	// reading till enter or new line is present
	input, _ := reader.ReadString('\n')
	// converting the input to type int
	parsedInput, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 32)
	fmt.Printf("%T and %T \n", parsedInput, int(parsedInput))

	isArmstrong(int(parsedInput))
}

func isArmstrong(val int) {
	sum := 0
	valCopy := val

	for valCopy > 0 {
		digit := valCopy % 10
		sum = sum + digit*digit*digit
		valCopy = valCopy / 10
	}

	if sum == val {
		fmt.Println("The number is an Armstrong number")
	} else {
		fmt.Println("The number is NOT an Armstrong number")
	}
}
