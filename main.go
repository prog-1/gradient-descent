package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	numberOfPoints     = 20
	pointMin, pointMax = 30, 70 //point distribution
	lineMin, lineMax   = 0, 100 //line lenght
)

func main() {

	//####################### Linear Regression #########################

	//Generating random points
	px := make([]float64, numberOfPoints)
	py := make([]float64, numberOfPoints)
	for i := 0; i < numberOfPoints; i++ {
		px[i] = (rand.Float64()*(pointMax-pointMin) + pointMin)
		py[i] = (rand.Float64()*(pointMax-pointMin) + pointMin)
	}

	//####################### Ebiten ####################################

	//Window
	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Linear Regression")

	//App instance
	a := NewApp(sW, sH)

	//
	go func() {
		a.linearRegression(px, py)
	}()

	//Running game
	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}

}
