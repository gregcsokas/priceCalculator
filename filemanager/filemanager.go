package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	fp, err := os.Open(path)

	if err != nil {
		return nil, errors.New("an error occurred while opening prices.txt")
	}

	scanner := bufio.NewScanner(fp)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fp.Close()
		return nil, errors.New("an error occurred while reading prices.txt")
	}

	fp.Close()

	return lines, nil
}
