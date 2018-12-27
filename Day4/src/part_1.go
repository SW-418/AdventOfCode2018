package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "input.txt"
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(file))
	}
}
