package main

import (
	"errors"
	"math"
)

func pearson(numbers []float64) (float64, error) {
	n := float64(len(numbers))
	var sumY float64
	for _, number := range numbers {
		sumY += number
	}
	meanY := sumY / n
	meanX := (n - 1) / 2

	var sumXY float64
	var sumXX float64
	var sumYY float64

	for i, number := range numbers {
		x := float64(i)
		dx := x - meanX
		dy := number - meanY
		sumXY += (dx * dy)
		sumXX += (dx * dx)
		sumYY += (dy * dy)

	}
	if sumYY == 0 {
		return 0, errors.New("Pearson is undefined for constant data")
	}

	denominator := math.Sqrt(sumXX * sumYY)
	r := sumXY / denominator

	return r, nil

}
