package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	X      float64
	F      func(float64) float64
	FDeriv func(float64) float64
	Img    <-chan *image.RGBA

	img *ebiten.Image
}

func (app *App) Update() error { return nil }

func (app *App) Draw(screen *ebiten.Image) {
	select {
	case img := <-app.Img:
		app.img = ebiten.NewImageFromImage(img)
	default:
	}
	if app.img != nil {
		screen.DrawImage(app.img, nil)
	}
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
