package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// 
	// TYPE 1
	// 

	// for d:=0; d<len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// 
	// TYPE 2
	// 
	
	// ALERT: i has index not the value
	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v \n", index, day)
	// }
	
	// for _, day := range days {
	// 	fmt.Printf("value is %v \n", day)
	// }

	rougeValue := 1
	
	for rougeValue < 10 {

		if rougeValue == 8 {
			break
		} else if rougeValue == 7 {
			rougeValue++
			continue
		}

		if rougeValue == 2 {
			goto lco
		}

		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

	// this is a goto
	lco:
		fmt.Println("Jumping at LearnCodeOnline.in")

}