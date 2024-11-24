package filemanager

import (
	"bufio"
	"encoding/json"
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

func WriteJSON(path string, data any) error {
	fp, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create file")
	}

	err = json.NewEncoder(fp).Encode(data)

	if err != nil {
		fp.Close()
		return errors.New("failed to encode file")
	}

	fp.Close()
	return nil
}
