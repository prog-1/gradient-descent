package main

import (
	"image"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func f(x float64) float64 { return -(1+x)*math.Pow(10, (-2)*x) + 8 }

func derivative(f func(float64) float64, x, e float64) float64 {
	return (f(x+e) - f(x-e)) / (2 * e)
}

func sign(x float64) float64 {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func gradientlength(a float64) float64 {
	return a / math.Sqrt(a*a)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gradient descent")

	img := make(chan *image.RGBA, 1)
	go func() {
		e := 1e-6
		p := Plot(-5, 0, 0.1, f)
		x := 0.0
		img <- p(x)
		lr := 1e-3
		for derivative(f, x, e) > 1e-6 {
			time.Sleep(30 * time.Millisecond)
			// x -= lr * sign(derivative(f, x, e))
			gradient := derivative(f, x, e)
			normalized := (1 / gradientlength(gradient)) * gradient
			x -= lr * sign(normalized)
			img <- p(x)
		}
	}()

	if err := ebiten.RunGame(&App{Img: img}); err != nil {
		log.Fatal(err)
	}
}
