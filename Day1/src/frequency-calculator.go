package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileName := "frequency.txt"
	fmt.Println("Henlo")

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

func ShouldAdd(s string) bool {

}
