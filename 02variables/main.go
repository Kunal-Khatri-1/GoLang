package main

import "fmt"

// var jwtToken := 300000 => not allowed to use walrus operator outside method
// var jwtToken = 300000 => allowed
// var jwtToken int = 300000 => allowed

const LoginToken string = "asdfhohds" // first letter capital means public variable, this is accessible by any other file into this folder

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n", smallVal)

	// five digits after the decimals are stored here
	var smallFloat float32 = 255.4564344346567453453452352
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	// more number of digits after decimals
	var largeFloat float64 = 255.4564344346567453453452352
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type %T \n", largeFloat)

	// default values and some aliases
	var anotherVariable int // stores zero
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T \n", anotherVariable)

	var anotherString string // stores ""
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type %T \n", anotherString)

	// implicit type/way to declare variable
	var website = "learncodeonline.in" // lexer decides the type of the variable if not mentioned
	fmt.Println(website)
	// website = 3	=> you can't change it to string

	// no var style
	numberOfUser := 300000.0 // walrus symbol
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)
}
