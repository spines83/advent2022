package utils

import (
	"bufio"
	"os"
)

func IterLines(path string) (<-chan string, error) {
	readFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	chn := make(chan string)
	go func() {
		for fileScanner.Scan() {
			chn <- fileScanner.Text()
		}
		readFile.Close()
		close(chn)
	}()

	return chn, nil
}

// Quick and dirty utility to count lines in a file
func CountLines(inputFile string) int {
	file, _ := os.Open(inputFile)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}
