package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readData(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []float64
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		number, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return numbers, nil

}
