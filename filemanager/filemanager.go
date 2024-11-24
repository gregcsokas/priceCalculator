package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		inputFilePath,
		outputFilePath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	fp, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("an error occurred while opening prices.txt")
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("an error occurred while reading prices.txt")
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	fp, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer fp.Close()

	time.Sleep(3 * time.Second)

	err = json.NewEncoder(fp).Encode(data)

	if err != nil {
		return errors.New("failed to encode file")
	}

	return nil
}
