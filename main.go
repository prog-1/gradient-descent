package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"gonum.org/v1/plot/plotter"
)

const (
	screenWidth, screenHeight                    = 720, 480
	randMin, randMax                             = -10, 10
	epochs, lr                                   = 5, 0.25
	plotMinX, plotMaxX, plotMinY, plotMaxY       = -10, 10, -50, 100 // Min and Max data values along both axis
	pointMinYOffset, pointMaxYOffset, pointCount = -20, 20, 10
)

// Function points are spawed along
func f(x float64) float64 {
	return 0.5*x + 2
}

// Approximating function(af) = Inference for 1 argument(x)
func af(x, a, b float64) float64 { return a*x + b }

// Runs model on all the input data
func inference(x []float64, w, b float64) (out []float64) {
	for _, v := range x {
		out = append(out, af(v, w, b))
	}
	return
}

func loss(labels, y []float64) float64 {
	var errSum float64
	for i := range labels {
		errSum += (y[i] - labels[i]) * (y[i] - labels[i])
	}
	return errSum / float64(len(labels)) // For the sake of making numbers smaller -> better percievable
}

func gradient(labels, y, x []float64) (dw, db float64) {
	// dw, db - Parial derivatives, w - weight, b - bias
	for i := 0; i < len(labels); i++ {
		dif := y[i] - labels[i]
		dw += dif * x[i]
		db += dif
	}
	n := float64(len(labels))
	dw *= 2 / n
	db *= 2 / n

	return
}

func train(epochs int, inputs, labels []float64) (w, b float64) {
	// randFloat64 := func() float64 {
	// 	return randMin + rand.Float64()*(randMax-randMin)
	// }
	// w, b = randFloat64(), randFloat64()
	w, b = 1, 0
	var dw, db float64
	for i := 0; i < epochs; i++ {
		dw, db = gradient(labels, inference(inputs, w, b), inputs)
		w -= dw * lr
		b -= db * lr
	}
	return
}

// Returns random points along f with random Yoffset
func randPoints(f func(float64) float64, minYoffset float64, maxYoffset float64, pointCount uint) (xs, labels []float64) {
	// 1. Getting random argument value X
	// 2. Getting function value(Y)
	// 3. Applying offset to Y
	for i := uint(0); i < pointCount; i++ {
		x := plotMinX + rand.Float64()*(plotMaxX-plotMinX) // Random argument within visible range
		yOffset := minYoffset + rand.Float64()*(maxYoffset-minYoffset)
		xs = append(xs, x)
		labels = append(labels, f(x)+yOffset)
	}
	return
}

// func main() {
// 	xs, labels := randPoints(f, pointMinYOffset, pointMaxYOffset, pointCount)
// 	w, b := train(epochs, labels, xs)
//
// 	img := make(chan *image.RGBA, 1)
// 	go func() {
// 		p := Plot(plotMinX, plotMaxX, )
// 	}
// 	if err := ebiten.RunGame(&App{Img: img}); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// var loss plotter.XYs
//
//	for i := 0; i < epochs; i++ {
//	  y := inference(inputs, w, b)
//	  loss = append(loss, plotter.XY{
//	    X: float64(i),
//	    Y: msl(labels, y),
//	  })
//	  lossLines, _ := plotter.NewLine(loss)
//	  if plotLoss {
//	    select {
//	    case img <- Plot(lossLines):
//	    default:
//	    }
//	  } else {
//	    const extra = (inputPointsMaxX - inputPointsMinX) / 10
//	    xs := []float64{inputPointsMinX - extra, inputPointsMaxX + extra}
//	    ys := inference(xs, w, b)
//	    resLine, _ := plotter.NewLine(plotter.XYs{{X: xs[0], Y: ys[0]}, {X: xs[1], Y: ys[1]}})
//	    img <- Plot(inputsScatter, resLine)
//	  }

func main() {
	inputs, labels := randPoints(f, pointMinYOffset, pointMaxYOffset, pointCount)
	var points plotter.XYs
	for i := 0; i < pointCount; i++ {
		points = append(points, plotter.XY{X: inputs[i], Y: labels[i]})
	}

	img := make(chan *image.RGBA, 1)
	pointsScatter, _ := plotter.NewScatter(points)
	fp := plotter.NewFunction(f) // f plot
	w, b := train(epochs, inputs, labels)
	fmt.Println(w, b)
	ap := plotter.NewFunction(func(x float64) float64 { return w*x + b }) // approximating function plot
	img <- Plot(pointsScatter, fp, ap)

	if err := ebiten.RunGame(&App{Img: img}); err != nil {
		log.Fatal(err)
	}
}
