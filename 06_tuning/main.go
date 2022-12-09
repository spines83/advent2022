package main

import (
	"advent2022/utils"
	"errors"
	"fmt"
)

func main() {
	lines, _ := utils.IterLines("input.txt")
	for currentLine := range lines {
		marker, _ := GetFirstMarker(currentLine, 14)
		fmt.Printf("%d\n", marker)
	}
}

func GetFirstMarker(inputLine string, length int) (int, error) {

	if len(inputLine) < length {
		return 0, errors.New("input line is too short")
	}

	strArr := []rune(inputLine)
	for i := length; i < len(strArr); i++ {

		// Check the last <length> characters to see if there are any dupes
		index := make(map[string]bool)
		success := true
		for j := i - length; j < i; j++ {
			character := string(strArr[j])
			if index[character] == true {
				success = false
				break
			}
			index[character] = true
		}

		// If there's no dupes, we've found our marker
		if success {
			return i, nil
		}
	}

	return 0, errors.New("Failed to discover a marker")
}
