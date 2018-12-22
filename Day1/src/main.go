package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ShouldAdd(s string) bool {
	if strings.Contains(s, "-") {
		return false
	}
	return true
}

func main() {
	fileName := "frequency.txt"
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	} else {
		unsplitString := string(file)
		splitString := strings.Split(unsplitString, "\n")
		fmt.Println(splitString)
		fmt.Println(" ")
		fmt.Println(splitString[0])

	}
}
