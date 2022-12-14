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
	// response is of type *http.Response
	fmt.Printf("Response is of type: %T \n", response)
	// caller's responsibility to close the connection
	defer response.Body.Close()

	// majority of reading done by ioutil
	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println(content)

}
