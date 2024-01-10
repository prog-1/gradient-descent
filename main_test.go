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
		{name: "diff.135 lpp. N18 a)", f: func(x float64) float64 { return 2*x*x - 3*x }, x: 3, want: 9},
		{name: "diff.135 lpp. N18 b)", f: func(x float64) float64 { return x*x - x + 2 }, x: 0, want: -1},
		{name: "diff.135 lpp. N18 c)", f: func(x float64) float64 { return 4 - 5*x - x*x }, x: 1, want: -7},
		{"diff.135 lpp. N18 d)", func(x float64) float64 { return -3 / x }, 3, 0.3333333}, //1/3
		{"diff.135 lpp. N18 e)", func(x float64) float64 { return 4 / x * x }, -2, 0},
		{"diff.135 lpp. N18 f)", func(x float64) float64 { return 4/x + 1 }, -2, -1},
		{"diff.135 lpp. N18 g)", func(x float64) float64 { return 6/2 - x }, -1, -1},
		{"diff.135 lpp. N18 h)", func(x float64) float64 { return 1 / math.Sqrt(x) }, 4, -0.062499},
		{"diff.135 lpp. N18 i)", func(x float64) float64 { return 2 * math.Sqrt(1-x) }, 0, -1},
		{"diff.135 lpp. N18 j)", func(x float64) float64 { return math.Sqrt(1 + 2*x) }, 4, 0.3333333},
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

func TestGradientDescend(t *testing.T) {
	for _, tc := range []struct {
		name  string
		f     func(float64) float64
		x0    float64
		alpha float64
		want  float64
	}{
		{name: "constant", f: func(x float64) float64 { return 5 }, x0: 0.0, alpha: 0.01, want: 0.0},
		{name: "simple square", f: func(x float64) float64 { return x * x }, x0: 2.0, alpha: 0.01, want: 0.0},
		{name: "not so simple square", f: func(x float64) float64 { return x*x + 5*x - 3 }, x0: 1.0, alpha: 0.01, want: -2.5},
		{name: "diff.135 lpp. N18 a)", f: func(x float64) float64 { return 2*x*x - 3*x }, x0: 3, alpha: 0.01, want: 0.75},
		{name: "diff.135 lpp. N18 b)", f: func(x float64) float64 { return x*x - x + 2 }, x0: 0, alpha: 0.01, want: 0.5},
	} {
		t.Run("", func(t *testing.T) {
			got := GradientDescent(tc.f, tc.x0, tc.alpha)
			if math.Abs(got-tc.want) > 0.0001 {
				t.Errorf("Expected GradientDescent(%f, %f) to be %f, but got %f", tc.x0, tc.alpha, tc.want, got)
			}
		})
	}
}
