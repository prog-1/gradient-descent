package main

import "math/rand"

type point struct {
	x, y float64
}

func NewRandomPoints(num int) []point {
	points := make([]point, num)
	for i := 0; i < num; i++ {
		points[i] = point{rand.Float64() * 300, rand.Float64() * 300}
	}
	return points
}
