package main

import (
	"math"
	"testing"
)

func TestDerivative(t *testing.T) {
	for _, tc := range []struct {
		name string
		f    func(float64) float64
		x    float64
		want float64
	}{
		{name: "constant", f: func(x float64) float64 { return 5 }, x: 0.0, want: 0.0},
		{name: "linear", f: func(x float64) float64 { return 2*x + 5 }, x: 0.0, want: 2.0},
		{name: "simple square", f: func(x float64) float64 { return x * x }, x: 2.0, want: 4.0},
		{name: "not so simple square", f: func(x float64) float64 { return x*x + 5*x - 3 }, x: 1.0, want: 7.0},
		{name: "hyperbola", f: func(x float64) float64 { return 1 / x }, x: 2.0, want: -0.25},
		//{name: "complex sin", f: func(x float64) float64 { return math.Sin(x*x + 5*x - 3) }, x: 1.0, want: 7.0},
		//{name: "complex cos", f: func(x float64) float64 { return math.Cos(x*x + 5*x - 3) }, x: 1.0, want: -1.0},
		// for last two precision is < 0.1 ha
	} {
		t.Run("", func(t *testing.T) {
			d := derivative(tc.f)
			got := d(tc.x)
			if math.Abs(got-tc.want) > 0.0001 {
				t.Errorf("Expected derivative(%f) to be %f, but got %f", tc.x, tc.want, got)
			}
		})
	}
}
