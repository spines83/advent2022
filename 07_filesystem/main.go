package main

import (
	"advent2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	// We'll use a stack to store each level of directory
	// that we're currently traversing
	var pwd utils.Stack

	// We need to map directory to total file size
	dirSize := make(map[string]int)

	// Iterate over the file
	lines, _ := utils.IterLines("input.txt")
	for currentLine := range lines {

		splitStr := strings.Split(currentLine, " ")
		// Handle a command
		if splitStr[0] == "$" {

			// Handle cd
			if splitStr[1] == "cd" {
				if splitStr[2] == ".." {
					_, _ = pwd.Pop()
				} else {
					pwd.Push(splitStr[2])
				}
			}

		} else {
			// By convention, we are in a "ls" command.

			// This strconv will filter out directory listings since
			// they are not relevant until we cd into them.
			fileSize, _ := strconv.Atoi(splitStr[0])

			// We need to create a key that captures the absolute path
			path := ""
			for _, dir := range pwd {
				path += dir
				path += "/"

				// Initialize key if it doesn't exist
				if _, ok := dirSize[path]; !ok {
					dirSize[path] = 0
				}
				dirSize[path] += fileSize
			}

		}
	}

	part1Size := 0
	for _, size := range dirSize {
		if size <= 100000 {
			part1Size += size
		}
	}

	part2Size := math.MaxInt
	unusedSpace := 70000000 - dirSize["//"]
	for _, size := range dirSize {
		if unusedSpace+size >= 30000000 && size < part2Size {
			part2Size = size
		}
	}

	fmt.Printf("%d\n", part1Size)
	fmt.Printf("%d\n", part2Size)
}
