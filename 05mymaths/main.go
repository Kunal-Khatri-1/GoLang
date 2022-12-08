package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)


func main(){
	fmt.Println("")

	var myNumberOne int = 2
	var myNumberTwo float64 = 4

	// this will give an error
	// fmt.Println("The sum is: ", myNumberOne + myNumberTwo)

	// Loosing precision here
	fmt.Println("The sum is: ", myNumberOne + int(myNumberTwo))

	// 
	// RANDOM NUMBERS => by math: weaker and crypto: strong
	// 

	// 
	// USING MATH
	// 

	// Don't have to use Math.rand, rand will work fine
	// rand.Seed(30)	// without this, rand.Intn will give 1 only. With different values of seed we get different random numbers
	// fmt.Println(rand.Intn(5))

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// 
	// USING CRYPTO
	// 

	// this gives different random numbers
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}