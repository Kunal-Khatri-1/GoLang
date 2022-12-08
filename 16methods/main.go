package main

import "fmt"

func main() {
	kunal := User{"Kunal", "kunal@go.dev", true, 19}
	kunal.GetStatus()
	kunal.NewMail()		// wont change Email of kunal
	fmt.Println("Email is: ", kunal.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

// This is method
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

// Copy is being passed instead of actual value
// POINTERS comes to the rescue here
func (u User) NewMail() {
	u.Email = "Test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}