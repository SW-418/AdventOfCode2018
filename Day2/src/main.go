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
		frequencyMap := make(map[int]int)
		duplicateFrequencyFound := false
		for index := 0; !duplicateFrequencyFound; index++ {
			stringConversion, err := strconv.Atoi(splitString[index])
			if err != nil {
				fmt.Println(err)
			}
			resultingFrequency += stringConversion
			if frequencyMap[resultingFrequency] != 0 {
				duplicateFrequencyFound = true
			} else {
				frequencyMap[resultingFrequency] = 1
			}
			if index == len(splitString)-1 {
				index = -1
			}
		}
		fmt.Println(resultingFrequency)
	}
}
