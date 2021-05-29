package parser

import (
	"bufio"
	"os"
)

func ParseLinksFromFile(pathFile string) ([]string, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}
