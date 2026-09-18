package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage : go run . data.txt")
		os.Exit(1)
	}
	numbers, err := readData(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading data:", err)
		os.Exit(1)
	}
	if len(numbers) < 2 {
		fmt.Fprintln(os.Stderr, "At least two numbers are required.")
		os.Exit(1)
	}

	a, b := linearRegression(numbers)

	fmt.Println(numbers, a, b)
}
