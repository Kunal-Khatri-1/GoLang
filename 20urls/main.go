package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=asdnt234w"

func main() {
	fmt.Println("Welcome to handling URLs in Golang")
	fmt.Println(myUrl)

	// parsing the URL
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	// NOTICE for port => it is a method
	fmt.Println(result.Port())

	// stores all the query parameters in better format => key value format
	qparams := result.Query()
	fmt.Printf("The type of query params is %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])
	fmt.Println(qparams)

	for _, value := range qparams {
		fmt.Println("Param is: ", value)
	}

	// ---------------------------------------------------------------------------------------------

	// creating URL from parts

	// NOTICE have to pass the reference and not a copy
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "LCO.dev",
		Path:    "/tutcss",
		RawPath: "user=kunal",
	}

	// another way of converting to string
	// first one was string(<data>)
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
