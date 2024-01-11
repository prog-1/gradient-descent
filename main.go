package main

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	e = 1e-5
)

func f(x float64) float64 { return x*x*x*x + x*x + 5*x - 3 }

func df(x float64) float64 { return 4*x*x*x + 2*x + 5 }

// func Derivative(x float64, f func(x float64) float64) float64 {
// 	return (f(x+e) - f(x)) / e
// }

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gradient descent")

	img := make(chan *image.RGBA, 1)
	go func() {
		p := Plot(-10, 10, 0.1, f)
		x := 5.0
		img <- p(x, false)
		for i := 0; i < 5000; i++ {
			time.Sleep(30 * time.Millisecond)
			x -= df(x) * 0.001
			img <- p(x, false)
		}

	}()

	if err := ebiten.RunGame(&App{Img: img}); err != nil {
		log.Fatal(err)
	}
}

func Abs(a float64) float64 {
	if a < 0 {
		return a * -1
	}
	return a
}
