package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileName := "boxids.txt"
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	} else {
		doubleLetterCount := 0
		tripleLetterCount := 0
		boxIds := string(file)
		splitBoxIds := strings.Split(boxIds, "\n")
		for _, boxId := range splitBoxIds {
			hasDoubleLetter := false
			hasTripleLetter := false
			letterCountMap := make(map[string]int)
			for _, letter := range boxId {
				letterCountMap[string(letter)] += 1
			}
			for key := range letterCountMap {
				if letterCountMap[key] == 3 {
					hasTripleLetter = true
				} else if letterCountMap[key] == 2 {
					hasDoubleLetter = true
				}
			}
			if hasDoubleLetter {
				doubleLetterCount += 1
			}
			if hasTripleLetter {
				tripleLetterCount += 1
			}
		}
		fmt.Printf("Double Count: %d \n", doubleLetterCount)
		fmt.Printf("Triple Count: %d \n", tripleLetterCount)
		checksum := doubleLetterCount * tripleLetterCount
		fmt.Printf("Calculated Checksum: %d \n", checksum)
	}
}
