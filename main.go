package main

import (
	"image"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 720, 480
	lr                        = 0.15 // Learning rate
)

func f(x float64) float64  { return math.Cos(x) }
func df(x float64) float64 { return -math.Sin(x) }

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gradient descent")

	img := make(chan *image.RGBA, 1)
	go func() {
		p := Plot(-10, 5, 0.1, f)
		x := .25
		img <- p(x)
		for i := 0; df(x) != 0; i++ {
			time.Sleep(1000 * time.Millisecond)
			x -= sign(df(x)) * 1 / math.Min(float64(i)+1, 100)
			img <- p(x)
		}
	}()

	if err := ebiten.RunGame(&App{Img: img}); err != nil {
		log.Fatal(err)
	}
}

func sign(a float64) float64 {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}
