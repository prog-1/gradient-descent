package main

import (
	"fmt"
)

func derivative(f func(float64) float64, x, e float64) float64 {
	return (f(x+e) - f(x-e)) / (e)
}

func gradient(f func(x float64) float64, a, x, e float64) float64 {
	for derivative(f, x, e) > 1e-6 {
		x -= a * derivative(f, x, e)
	}
	return x
}

func main() {
	e := 1e-6
	f := func(x float64) float64 { return x*x*x*x - 4*x*x*x + 2*x*x + 10*x + 12 }
	min := gradient(f, 0.01, 1, e)
	fmt.Println(min)
}
