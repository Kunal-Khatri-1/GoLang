package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Encode map data into a JSON string")

	mapping := map[string]string{
		"Name": "Kunal",
		"Age":  "19",
	}

	bytesJsonMapping, err := json.Marshal(mapping)
	checkNilError(err)
	jsonMapping := string(bytesJsonMapping)

	println(jsonMapping)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
