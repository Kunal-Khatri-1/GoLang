package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bubble Sort")

	unsortedSlice := []int{3, 2, 1}
	sortedSlice := bubbleSort(unsortedSlice)

	fmt.Println(sortedSlice)
}

func bubbleSort(slice []int) []int {
	sorted := false
	n := 1
	for !sorted {
		// if flow control never enter the loop => slice has been sorted
		sorted = true
		for i := 0; i < len(slice)-n; i++ {
			sorted = false
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
		n++
	}

	return slice
}
