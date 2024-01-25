package main

import (
	"image"
	"log"

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
	p := NewRandomPoints(10)
	l := NewLine()
	err := l.Train(p, 0.001, 5000)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gradient descent")

	img := make(chan *image.RGBA, 1)
	go func() {
		p := Plot(-30, 30, 0.1, f)
		x := 30.0
		img <- p(x, false)
		for i := 0; i < 500000; i++ {
			// time.Sleep( * time.Millisecond)
			x -= df(x) * 1e-6
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
