package main

import (
	"bufio"
	"fmt"
	"math"
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
	lineNumber := 0

	var numbers []float64

	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()
		line = strings.TrimSpace(line)

		number, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, fmt.Errorf("invalid number on line %d", lineNumber)
		}
		if math.IsNaN(number) || math.IsInf(number, 0) {
			return nil, fmt.Errorf("invalid number on line %d: %s", lineNumber, line)

		}
		numbers = append(numbers, number)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return numbers, nil

}
