package main

import "fmt"

func main() {
	fmt.Println("Merge sort")

	unsortedSlice := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sortedSlice := mergeSort(unsortedSlice)

	fmt.Println(sortedSlice)
}

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	leftPart := mergeSort(slice[:len(slice)/2])
	rightPart := mergeSort(slice[len(slice)/2:])
	return merge(leftPart, rightPart)
}

func merge(leftSlice []int, rightSlice []int) []int {
	returnSlice := []int{}
	i := 0
	j := 0
	for i < len(leftSlice) && j < len(rightSlice) {
		if leftSlice[i] < rightSlice[j] {
			returnSlice = append(returnSlice, leftSlice[i])
			i++
		} else {
			returnSlice = append(returnSlice, rightSlice[j])
			j++
		}
	}
	for ; i < len(leftSlice); i++ {
		returnSlice = append(returnSlice, leftSlice[i])
	}
	for ; j < len(rightSlice); j++ {
		returnSlice = append(returnSlice, rightSlice[j])
	}
	return returnSlice
}

// package main

// import "fmt"

// func main() {
// 	unsortedSlice := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
// 	sortedSlice := mergeSort(unsortedSlice)

// 	fmt.Println(sortedSlice)
// 	// sorted = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// }

// func mergeSort(slice []int) []int {
// 	if len(slice) < 2 {
// 		return slice
// 	}
// 	leftPart := mergeSort(slice[:len(slice)/2])
// 	rightPart := mergeSort(slice[len(slice)/2:])
// 	return merge(leftPart, rightPart)
// }

// func merge(leftSlice []int, rightSlice []int) []int {
// 	returnSlice := []int{}
// 	i := 0
// 	j := 0
// 	for i < len(leftSlice) && j < len(rightSlice) {
// 		if leftSlice[i] < rightSlice[j] {
// 			returnSlice = append(returnSlice, leftSlice[i])
// 			i++
// 		} else {
// 			returnSlice = append(returnSlice, rightSlice[j])
// 			j++
// 		}
// 	}
// 	for ; i < len(leftSlice); i++ {
// 		returnSlice = append(returnSlice, leftSlice[i])
// 	}
// 	for ; j < len(rightSlice); j++ {
// 		returnSlice = append(returnSlice, rightSlice[j])
// 	}
// 	return returnSlice
// }
