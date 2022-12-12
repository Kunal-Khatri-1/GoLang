package main

import "fmt"

func main() {
	fmt.Println("Interface with type embedding and method overrding in Golang")

	cat := catStruct{task: "meow"}
	dog := dogStruct{task: "woof"}

	mammals := mammalStruct{mammal: []mammalInterface{cat, dog}}

	for _, value := range mammals.mammal {
		value.doTask()
	}
}

type catStruct struct {
	task string
}

func (c catStruct) doTask() {
	fmt.Println(c.task)
}

type dogStruct struct {
	task string
}

func (d dogStruct) doTask() {
	fmt.Println(d.task)
}

type mammalInterface interface {
	doTask()
}

type mammalStruct struct {
	mammal []mammalInterface
}
