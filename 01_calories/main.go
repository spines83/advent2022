package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"advent2022/utils"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	currentCalorieCounter := 0
	totalCaloriesByElf := []int{}

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		if currentLine == "" {
			totalCaloriesByElf = append(totalCaloriesByElf, currentCalorieCounter)
			currentCalorieCounter = 0
		} else {
			intVar, err := strconv.Atoi(currentLine)
			if err != nil {
				fmt.Println(err)
			}
			currentCalorieCounter += intVar
		}
	}

	readFile.Close()

	// One final check to make sure the last elf isn't holding the most
	totalCaloriesByElf = append(totalCaloriesByElf, currentCalorieCounter)
	sort.Ints(totalCaloriesByElf)

	top1 := totalCaloriesByElf[len(totalCaloriesByElf)-1]
	top3 := totalCaloriesByElf[len(totalCaloriesByElf)-3:]

	fmt.Printf("top 1: %d, top 3: %d\n", top1, utils.Sum(top3))
}
