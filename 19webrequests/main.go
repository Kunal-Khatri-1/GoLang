package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// response type is not type string or bytes
	fmt.Printf("Response is of type: %T \n", response)
	// caller's responsibility to close the connection
	defer response.Body.Close()

	ioutil

}
