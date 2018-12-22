package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileName := "frequency.txt"
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	} else {
		unsplitString := string(file)
		splitString := strings.Split(unsplitString, "\n")
		resultingFrequency := 0
		for index := 0; index < len(splitString); index++ {
			stringConversion, err := strconv.Atoi(splitString[index])
			if err != nil {
				fmt.Println(err)
			}
			resultingFrequency += stringConversion
		}
		fmt.Println(resultingFrequency)
	}
}
