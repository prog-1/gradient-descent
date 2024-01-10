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

func f(x float64) float64  { return 0.01*x*x*x*x + x*x + 5*x - 3 }
func df(x float64) float64 { return 2*x + 5 }

func Derivative(x float64, f func(x float64) float64) float64 {
	return (f(x+e) - f(x)) / e
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gradient descent")

	img := make(chan *image.RGBA, 1)
	go func() {
		p := Plot(-10, 10, 0.1, f)
		x := 10.0
		img <- p(x, false)
		for {
			time.Sleep(300 * time.Millisecond)
			x -= Derivative(x, f) * 0.3
			img <- p(x, false)
			if Abs(Derivative(x, f)) < e {
				img <- p(x, true)
				break
			}
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
