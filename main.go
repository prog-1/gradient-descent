package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func f(x float64) float64 { return math.Cos(x*x + 3*x + 97) } // FUCTION CONTROL

// TODO: use govaluate to get function as user input
// i can get string from user, parse it to function, but then what?

func derivative(f func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		dx := 0.00000000005          // dx -> 0
		return (f(x+dx) - f(x)) / dx //literally todays math lesson.
	}
}
func GradientDescent(f func(float64) float64, x0, alpha float64) float64 {
	x := x0
	var stop bool
	for !stop {
		x -= alpha * derivative(f)(x)
		if derivative(f)(x) == 0 || math.Abs(derivative(f)(x)) < 0.00001 {
			stop = true
		}
	}
	return x
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gradient descent")

	fmt.Println(GradientDescent(f, -math.Pi/6, 0.01)) //CONTROL X0 AND ALPHA HERE FOR PRINTLN OUTPUT

	img := make(chan *image.RGBA, 1)
	go func() {
		p := Plot(-math.Pi, math.Pi/3, 0.01, f) //CONTROL RANGE AND STEP HERE
		x := -math.Pi / 6                       //CONTROL X0 HERE FOR GRAPHICS
		img <- p(x)
		var stop bool
		for !stop {
			time.Sleep(30 * time.Millisecond)
			x -= 0.1 * derivative(f)(x) // CONTROL ALPHA HERE FOR GRAPHICS
			// if you put alpha = 0.9, you wil get badly animated film about snake :)
			if derivative(f)(x) == 0 || math.Abs(derivative(f)(x)) < 0.00001 {
				stop = true
			}
			img <- p(x)
		}
	}()

	if err := ebiten.RunGame(&App{Img: img}); err != nil {
		log.Fatal(err)
	}
}
