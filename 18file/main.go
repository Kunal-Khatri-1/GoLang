package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcofile.txt")

	checkNilErr(err)

	// io is used for writing
	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./mylcofile.txt")
}

func readFile(fileName string) {
	// the file is not being read in string format
	// it is read in bytes format
	dataByte, err := ioutil.ReadFile(fileName)

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", dataByte)
	fmt.Println("Text data inside the file is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		// panic will stop the execution
		panic(err)
	}
}
