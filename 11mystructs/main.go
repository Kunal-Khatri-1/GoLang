package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or the concept of parent

	kunal := User{"Kunal", "Kunal@go.dev", true, 19}
	fmt.Println(kunal)
	// using %+v give much more details compared to %v
	fmt.Printf("Kunal details are: %+v \n", kunal)
	fmt.Printf("Name is %v and email is %v \n", kunal.Name, kunal.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}