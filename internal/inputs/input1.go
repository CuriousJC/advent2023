package inputs

import (
	"bufio"
	"os"
	"path/filepath"
)

func Input1() ([]string, error) {

	lines, err := readLinesFromFile("static/input1.txt")
	if err != nil {
		panic(err)
	}

	return lines, nil

}

func readLinesFromFile(filename string) ([]string, error) {
	// Construct the absolute file path based on the relative path to the current directory
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
