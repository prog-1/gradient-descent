package main

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func f(x float64) float64  { return x*x + 5*x - 3 }
func df(x float64) float64 { return 2*x + 5 }

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gradient descent")

	img := make(chan *image.RGBA, 1)
	go func() {
		p := Plot(-5, 0, 0.1, f)
		x := 0.0
		img <- p(x)
		for i := 0; i < 50; i++ {
			time.Sleep(30 * time.Millisecond)
			x -= df(x) * 0.1
			img <- p(x)
		}
	}()

	app := &App{
		X:      0,
		F:      f,
		FDeriv: df,
		Img:    img,
	}
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
