package main

import (
	"math"
	"testing"
)

func TestSolveSecant(t *testing.T) {
	for _, tc := range []struct {
		f         func(float64) float64
		a, x, eps float64
		want      float64
	}{
		{f: func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, a: 0.01, x: 2, eps: 1e-3, want: 0.816},
		{f: func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 }, a: 0.01, x: 1, eps: 1, want: 2},
		{f: func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, a: 0.09, x: 10, eps: 1e-5, want: 7.65262},
		{f: func(x float64) float64 { return (1+x)*math.Pow(10, 3*x) - 5 }, a: 0.01, x: 0.01, eps: 1e-3, want: -1.145},
		{f: func(x float64) float64 { return -(1+x)*math.Pow(10, (-2)*x) + 8 }, a: 0.001, x: 2, eps: 1e-3, want: -0.783},
	} {
		got := gradient(tc.f, tc.a, tc.x, tc.eps)
		if math.Abs(got-tc.want) > tc.eps {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}

	}
}
