package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	// compulsory to provide how long the array is gonna be
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	// index 2 will store ""
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", len(vegList))

	// there is not sorting or something crazy for arrays in golang
}