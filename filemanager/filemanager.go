package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
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

	scanner := bufio.NewScanner(fp)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		err := fp.Close()
		if err != nil {
			return nil, err
		}
		return nil, errors.New("an error occurred while reading prices.txt")
	}

	err = fp.Close()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	fp, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	err = json.NewEncoder(fp).Encode(data)

	if err != nil {
		err := fp.Close()
		if err != nil {
			return err
		}
		return errors.New("failed to encode file")
	}

	err = fp.Close()
	if err != nil {
		return err
	}
	return nil
}
