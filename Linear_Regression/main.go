package main

import (
	"fmt"
)

type Line struct {
	k float64
	b float64
}

func (l *Line) Training(x, y []float64, learningRate float64, epochs int) {
	m := float64(len(x))

	for epoch := 1; epoch <= epochs; epoch++ {

		// Making calculations
		calculations := make([]float64, len(x))
		for i := range x {
			calculations[i] = l.calculate(x[i])
		}

		// Calculating errors
		errors := make([]float64, len(y))
		for i := range y {
			errors[i] = calculations[i] - y[i]
		}

		l.k -= learningRate * (1 / m) * twoSliceSum(errors, x)
		l.b -= learningRate * (1 / m) * sliceSum(errors)

		loss := average(errors)
		fmt.Printf("Epoch %d, Cost: %f\n", epoch, loss)
	}
}

func (l *Line) calculate(x float64) float64 {
	return l.k*x + l.b
}

func twoSliceSum(slice1, slice2 []float64) float64 {
	result := 0.0
	for i := range slice1 {
		result += slice1[i] * slice2[i]
	}
	return result
}

func sliceSum(slice1 []float64) float64 {
	result := 0.0
	for i := range slice1 {
		result += slice1[i]
	}
	return result
}

func average(s []float64) float64 {
	return sliceSum(s) / float64(len(s))
}

func main() {
	// Points
	x := []float64{0, -0.5, 0.5, 1, 2.5}
	y := []float64{1, 0, 2, 3, 6}

	//Line initialization
	lr := Line{}

	// Training
	learningRate := 0.01
	epochs := 1000

	lr.Training(x, y, learningRate, epochs)

	fmt.Printf("Trained Coefficients: Slope=%.2f, Intercept=%.2f\n", lr.k, lr.b)

	newX := 6.0
	regression := lr.calculate(newX)
	fmt.Printf("Regression for X=%.2f: %.2f\n", newX, regression)
}
