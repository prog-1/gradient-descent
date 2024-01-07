package main

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func f(x float64) float64      { return x*x + 5*x - 3 }
func fDeriv(x float64) float64 { return 2*x + 5 }

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gradient descent")

	img := make(chan *image.RGBA, 1)
	go func() {
		p := Plot(-10, -10, 10, 100, f)
		x := 5.0
		img <- p(x)
		for i := 0; i < 200; i++ {
			time.Sleep(20 * time.Millisecond)
			x -= fDeriv(x) * 0.1
			img <- p(x)
		}
	}()

	app := &App{
		X:      0,
		F:      f,
		FDeriv: fDeriv,
		Img:    img,
	}
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
