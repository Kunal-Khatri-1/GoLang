package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Regular Expression to extract domain from URL")
	givenURL := "https://mastersunion.instructure.com/courses/283/assignments/1899"
	givenURL = "https://www.google.com/somepage/something/etc"
	regex := regexp.MustCompile(`\/\/(www.)?([A-Za-z]*\.)*[a-z]*`)

	fmt.Println(regex.MatchString(givenURL))

	// allMatches := regex.FindAllString(givenURL, -1)
	match := regex.FindString(givenURL)
	domain := match[2:]
	fmt.Println(domain)
}
