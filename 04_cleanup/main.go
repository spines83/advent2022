package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalContained := 0
	totalOverlap := 0
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		pair := strings.Split(currentLine, ",")
		elf1, elf2 := pair[0], pair[1]
		elf1_split := strings.Split(elf1, "-")
		elf2_split := strings.Split(elf2, "-")
		elf1_left, _ := strconv.Atoi(elf1_split[0])
		elf1_right, _ := strconv.Atoi(elf1_split[1])
		elf2_left, _ := strconv.Atoi(elf2_split[0])
		elf2_right, _ := strconv.Atoi(elf2_split[1])

		if IsContained(elf1_left, elf1_right, elf2_left, elf2_right) {
			totalContained += 1
		}

		if AnyOverlap(elf1_left, elf1_right, elf2_left, elf2_right) {
			totalOverlap += 1
		}

	}

	readFile.Close()
	fmt.Printf("contained: %d, overlap: %d\n", totalContained, totalOverlap)
}

func IsContained(e1_left, e1_right, e2_left, e2_right int) bool {
	if e1_left <= e2_left && e2_right <= e1_right {
		return true
	} else if e2_left <= e1_left && e1_right <= e2_right {
		return true
	} else {
		return false
	}
}

func AnyOverlap(e1_left, e1_right, e2_left, e2_right int) bool {
	if e1_left <= e2_left && e2_left <= e1_right {
		return true
	} else if e1_left <= e2_right && e2_right <= e1_right {
		return true
	} else if e2_left <= e1_left && e1_left <= e2_right {
		return true
	} else if e2_left <= e1_right && e1_right <= e2_right {
		return true
	} else {
		return false
	}
}
