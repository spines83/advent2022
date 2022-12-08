package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	scores := buildScoreMap()

	totalScore := 0
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		l := strings.Split(currentLine, " ")
		play, response := l[0], l[1]
		score := scores[play][response]
		totalScore += score
		fmt.Printf("%s + %s = %d\n", play, response, score)
	}

	readFile.Close()
	fmt.Printf("%d\n", totalScore)
}

func buildScoreMap() map[string]map[string]int {

	// fileContent, err := os.Open("scores.json")
	fileContent, err := os.Open("scores_p2.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var result map[string]map[string]int
	json.Unmarshal(byteResult, &result)
	return result
}
