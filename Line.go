package main

import (
	"fmt"
	"log"
	"math"
)

type line struct {
	k float64
	b float64
}

func (l *line) y(x float64) float64 {
	return l.k*x + l.b
}

func NewLine() *line {
	return &line{0, 0}
}

func (l *line) Train(points []point, lr float64, epochs uint) error {
	for i := uint(0); i < epochs; i++ {
		y := make([]float64, len(points))
		for j := range points {
			y[j] = l.y(points[j].x)
		}

		var sum1, sum2 float64
		for j := range points {
			sum1 += y[j] - points[j].y
			sum2 += y[j] - points[j].y*points[j].x
		}

		l.k -= lr * (2 / float64(len(points))) * sum2
		l.b -= lr * (2 / float64(len(points))) * sum1
		loss := sum1 / float64(len(points))
		if math.IsNaN(loss) || math.IsInf(loss, 0) {
			return fmt.Errorf("training unsuccessful")
		}
		log.Printf("Epoch: %d/%d, Loss %f\n", i, epochs, loss)
	}
	return nil
}
