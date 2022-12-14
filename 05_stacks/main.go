package main

import (
	"advent2022/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	// We use a stack to track the input
	var inputStack utils.Stack

	// We also need an array of stacks to track the individual crate-stacks
	// This will get initialized later once we figure out how large the input is
	var stackArray []utils.Stack

	lines, _ := utils.IterLines("full_input.txt")
	for currentLine := range lines {
		if IsABoxLine(currentLine) {
			// Defer processing the inital crates until we've scanned them all
			// Use a stack so we can put them back in the proper order
			inputStack.Push(currentLine)
		} else if IsTheNumberLine(currentLine) {
			// Kinda by convention, but we use this as the trigger to allocate
			// the stack of stacks.
			stackSize := len(currentLine)/4 + 1
			stackArray = make([]utils.Stack, stackSize)
		} else if len(currentLine) == 0 {
			// Leaning more into assumptions, the blank line is the trigger to
			// build out the initial data structure of crates
			buildInitialStack(stackArray, inputStack)
		} else {
			// Start executing instructions
			// executeInstruction(stackArray, currentLine)
			crateMover9001(stackArray, currentLine)
		}
	}
	// Once all of the instructions are ran, the top element of
	// each stack is the solution
	for i := 0; i < len(stackArray); i++ {
		topCrate, _ := stackArray[i].Peek()
		fmt.Printf("%s", topCrate)
	}
}

func executeInstruction(stackArray []utils.Stack, currentLine string) {
	// https://www.golangprograms.com/regular-expression-to-extract-numbers-from-a-string-in-golang.html
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	matches := re.FindAllString(currentLine, -1)
	if len(matches) != 3 {
		log.Fatal("whoops")
	}
	count, _ := strconv.Atoi(matches[0])
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])

	for i := 0; i < count; i++ {
		// Take one down, pass it around~
		if !stackArray[from-1].IsEmpty() {
			crate, _ := stackArray[from-1].Pop()
			stackArray[to-1].Push(crate)
		}
	}
}

func crateMover9001(stackArray []utils.Stack, currentLine string) {
	// https://www.golangprograms.com/regular-expression-to-extract-numbers-from-a-string-in-golang.html
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	matches := re.FindAllString(currentLine, -1)
	if len(matches) != 3 {
		log.Fatal("whoops")
	}
	count, _ := strconv.Atoi(matches[0])
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])

	var tmp utils.Stack
	for i := 0; i < count; i++ {
		// Take one down, pass it around~
		if !stackArray[from-1].IsEmpty() {
			crate, _ := stackArray[from-1].Pop()
			tmp.Push(crate)
		}
	}
	for len(tmp) > 0 {
		crate, _ := tmp.Pop()
		stackArray[to-1].Push(crate)
	}
}

func buildInitialStack(stackArray []utils.Stack, inputStack utils.Stack) {
	for len(inputStack) > 0 {
		inputLine, _ := inputStack.Pop()
		strArr := []rune(inputLine)
		// Taking advantage of every 4 characters being a potential letter
		for i := 1; i < len(strArr); i += 4 {
			crate := string(strArr[i])
			if crate != " " {
				stackArray[i/4].Push(crate)
			}
		}
	}
}

func IsABoxLine(line string) bool {
	return strings.Contains(line, "[")
}

func IsTheNumberLine(line string) bool {
	return len(line) > 1 && unicode.IsDigit(rune(line[len(line)-2]))
}
