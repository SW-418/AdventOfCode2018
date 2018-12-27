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
		boxIds := string(file)
		splitBoxIds := strings.Split(boxIds, "\n")
		matchingStrings := make(map[string]string)
		for outerIndex := 0; outerIndex < len(splitBoxIds); outerIndex++ {
			for innerIndex := 1; innerIndex < len(splitBoxIds)-1; innerIndex++ {
				lettersNotMatching := 0
				for letterIndex := 0; letterIndex < len(splitBoxIds[outerIndex]) && lettersNotMatching < 2; letterIndex++ {
					outerLetter := string(splitBoxIds[outerIndex][letterIndex])
					innerLetter := string(splitBoxIds[innerIndex][letterIndex])
					if outerLetter != innerLetter {
						lettersNotMatching++
					}
				}
				if lettersNotMatching < 2 && lettersNotMatching > 0 {
					matchingStrings[splitBoxIds[outerIndex]] = splitBoxIds[innerIndex]
				}
			}
		}
		fmt.Println(matchingStrings)
	}
}
