package main

import (
	"fmt"
)

func main() {
	fmt.Println("Create and Print Multi Dimensional Slice in Golang")
	multiSlice := [][]string{{"kunal", "19"}, {"sagar", "18"}}

	for _, slice := range multiSlice {
		for _, val := range slice {
			fmt.Println("The value is: ", val)
		}
	}
}
