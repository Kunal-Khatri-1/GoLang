package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// 01-02-2006 => standard for formatting
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.November, 10, 00, 00, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}