package main

import "fmt"

// Walrus operator will not work here outside the function/method
// jwtToken := 300000
// These are allowed here
var jwtToken = 300000
var jwtToken2 int = 300000

// 
// CONSTANTS => cannot be changed
// 

// NOTICE: capital first letter => PUBLIC VARIABLE
// This LoginToken is now accessible in any other file in the program
const LoginToken string = "asiufgi"

func main() {
	// declaring a variable and not using it will raise an error
	var username string = "Kunal"
	fmt.Println(username)
	// %T is the placeholder
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// 256 will raise an error
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// Only 5 values after the decimals and rest of the values are being ignored
	var smallFloat float32 = 255.21828182818281828128182818281
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	// more precision
	var bigFloat float64 = 255.21828182818281828128182818281
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type %T \n", bigFloat)

	// 
	// DEFAULT VALUES AND ALIASES
	// 

	// No garbage values, it has 0 by default
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// "" by default
	var anotherVariableString string
	fmt.Println(anotherVariableString)
	fmt.Printf("Variable is of type: %T \n", anotherVariableString)
	
	// 
	// IMPLICIT TYPE
	// 

	// no errors if we did not mention the type
	// Lexer will automatically assign the type based on the value passed
	// it won't allow to change the value to other type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type %T \n", website)
	// website = 987654 	// this gives an error

	// 
	// No var Style
	// 

	// Walrus Operator
	// inside a method Walrus operator is allowed
	// outside the method Walrus operator is not allowed
	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type %T \n", numberOfUsers)

	// Accessing Global LoginToken
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)
}