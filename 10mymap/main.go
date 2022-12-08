package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	// add key-value pair
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	// delete key-value pair
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interseting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
	
	for _, value := range languages {
		fmt.Printf("value is %v \n", value)
	}
}