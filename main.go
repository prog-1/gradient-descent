package main

import (
	"image"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

//FUNCTIONS

//func f(x float64) float64 { return x*x + 3*x - 5 }

func f(x float64) float64 { return math.Cos(x) }

// func f(x float64) float64 { return math.Sin(x) }
//func f(x float64) float64 { return 0.1*(x*x) - math.Sqrt(97)*x + 10 }

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Meh")

	img := make(chan *image.RGBA, 1)
	go func() {
		p := Plot(-10, 10, 0.1, f)
		x := 0.25
		img <- p(x)
		for i := 0; d(f, x) != 0; i++ {
			time.Sleep(500 * time.Millisecond)
			x -= d(f, x) * 0.9
			img <- p(x)
		}
	}()

	if err := ebiten.RunGame(&App{Img: img}); err != nil {
		log.Fatal(err)
	}
}

// Isolated one-dimensional gradient descent function (not used)
func GradientDescent(f func(float64) float64) float64 {
	x := 0.25
	for i := 0; d(f, x) != 0; i++ {
		x -= d(f, x) * 0.9
	}
	return x
}

// Derivative
func d(f func(float64) float64, x float64) float64 {
	h := 0.0001
	return (f(x+h) - f(x)) / h
}
