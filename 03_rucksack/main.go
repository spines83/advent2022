package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Part 1
	// totalScore := 0
	// for fileScanner.Scan() {
	// 	line := fileScanner.Text()
	// 	first := line[:len(line)/2]
	// 	second := line[len(line)/2:]
	// 	commonItem, err := FindCommonItem([]string{first, second})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	score := GetPriority(commonItem)
	// 	fmt.Printf("%s %s %c %d\n", first, second, commonItem, score)
	// 	totalScore += score
	// }

	// Part 2
	var lines []string
	counter := 0
	totalScore := 0
	for fileScanner.Scan() {
		counter += 1
		currentLine := fileScanner.Text()
		lines = append(lines, currentLine)
		if counter%3 == 0 {
			commonItem, err := FindCommonItem(lines)
			if err != nil {
				log.Fatal(err)
			}
			score := GetPriority(commonItem)
			totalScore += score
			lines = nil
		}
	}

	readFile.Close()
	fmt.Printf("%d\n", totalScore)
}

// This will find the common character (rune) between an arbitrary list of strings
func FindCommonItem(items []string) (rune, error) {

	// If there's only one item in the list, that's an issue
	if len(items) <= 1 {
		return -1, errors.New("This function should be called with more than one item")
	}

	// Index everything but the last string
	// For example, if the input was "acc", "cce", "ace" then indices would serialize to:
	// [
	// 	{
	// 		"a": true,
	// 		"c": true
	// 	},
	// 	{
	// 		"c": true,
	// 		"e": true
	// 	}
	// ]
	var indices []map[rune]bool
	for _, item := range items[:len(items)-1] {
		index := make(map[rune]bool)
		for _, character := range item {
			index[character] = true
		}
		indices = append(indices, index)
	}

	// For each character in the last item, we want to check
	// and make sure it's present in both indicies above.
	for _, c := range items[len(items)-1] {

		// The function assumes a return of true unless it hits
		// a condition that would set it to false
		outerTest := true

		for _, index := range indices {

			// If the last string isn't present in one of the indices
			// then we need to fail this character since it isn't our solution
			if index[c] == false {
				outerTest = false
			}
		}

		// If we've made it to this point without tripping a failure condition
		// then we've hit the one character that is present in all of the previous
		// strings. Go ahead and return it as the solution.
		if outerTest == true {
			return c, nil
		}
	}

	// There is no common item
	return -1, errors.New("There was no common item")
}

func GetPriority(x rune) int {
	if unicode.IsUpper(x) {
		return int(x) - 38
	} else {
		return int(x) - 96
	}
}
