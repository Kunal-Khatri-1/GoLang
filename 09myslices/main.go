package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	// we don't define how many values in slices
	// if values are defined => simple array
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	// this is not how to work with slices
	// fruitList[] = xyz 
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1: ])
	fmt.Println(fruitList)

	fruitList = append(fruitList[: 3])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1: 3], "DragonFruit", "Watermelon")
	fmt.Println(fruitList)

	// gives slice of integers and can store 4 values
	highScores := make([]int, 4)
	// this is is the default allocation of the memory
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777		// gives an error

	// as soon as we use append it is going to reallocate the memory to accomodate all the values
	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))


	// methods available for slices NOT arrays
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// 
	// HOW TO REMOVE A VALUE FROM SLICE BASED ON INDEX
	// 

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}