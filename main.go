package main

import "fmt"

type functiontype struct {
	k, b float64
}

type coord struct {
	x, y float64
}

func main() {
	points := []coord{{0, 1}, {-0.5, 0}, {0.5, 2}, {1, 3}, {2.5, 6}}
	learnRate := 0.01
	epochs := 1000

	train(points, learnRate, epochs)
}

func train(points []coord, learnRate float64, epochs int) {
	n := float64(len(points))
	var line functiontype

	for epoch := 1; epoch <= epochs; epoch++ {
		fxi := findFxi(points, line)
		er := findEr(points, fxi)

		var sumk, sumb float64
		for i, j := range er {
			sumk += j * points[i].x
			sumb += j
		}

		line.k += (2 / n) * sumk * learnRate
		line.b += (2 / n) * sumb * learnRate

		loss := average(er)

		if epoch%100 == 0 {
			fmt.Printf("Epoch %d, Loss: %f\n", epoch, loss)
		}
	}
}

func findFxi(points []coord, line functiontype) (fxi []float64) {
	for _, j := range points {
		fxi = append(fxi, line.k*j.x+line.b)
	}
	return fxi
}

func findEr(points []coord, fxi []float64) (er []float64) {
	for i, j := range points {
		er = append(er, j.y-fxi[i])
	}
	return er
}

func sliceSum(s []float64) (sum float64) {
	for _, j := range s {
		sum += j
	}
	return sum
}

func average(s []float64) float64 {
	return sliceSum(s) / float64(len(s))
}
