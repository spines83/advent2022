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
