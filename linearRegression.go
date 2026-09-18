package main

func linearRegression(numbers []float64) (float64, float64) {
	var sumY float64
	var sumXY float64
	var sumXX float64

	n := float64(len(numbers))
	for _, number := range numbers {
		sumY += number
	}

	meanY := sumY / n
	meanX := (n - 1) / 2
	for i, number := range numbers {
		x := float64(i)
		dx := x - meanX
		dy := number - meanY
		sumXY += dx * dy
		sumXX += dx * dx
	}
	a := sumXY / sumXX
	b := meanY - a*meanX
	return a, b

}
