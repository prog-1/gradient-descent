package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func Plot(xmin, ymin, xmax, ymax float64, f func(float64) float64) func(x float64) *image.RGBA {
	return func(x float64) *image.RGBA {
		fn := plotter.NewFunction(f)
		fn.Color = color.RGBA{B: 255, A: 255}

		pts := plotter.XYs{{X: x, Y: f(x)}}

		xScatter, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatalf("Failed to NewScatter: %v", err)
		}
		xScatter.Color = color.RGBA{R: 255, A: 255}

		labels, err := plotter.NewLabels(plotter.XYLabels{
			XYs:    pts,
			Labels: []string{fmt.Sprintf("x = %.2f", x)},
		})
		labels.Offset = vg.Point{X: -10, Y: 15}
		if err != nil {
			log.Fatalf("Failed to NewLabels: %v", err)
		}

		p := plot.New()
		p.Add(
			plotter.NewGrid(),
			fn,
			xScatter,
			labels,
		)
		p.Legend.Add("f(x)", fn)
		p.Legend.Add("x", xScatter)
		p.X.Label.Text = "X"
		p.Y.Label.Text = "Y"
		p.X.Min, p.Y.Min = xmin, ymin
		p.X.Max, p.Y.Max = xmax, ymax

		img := image.NewRGBA(image.Rect(0, 0, 640, 480))
		c := vgimg.NewWith(vgimg.UseImage(img))
		p.Draw(draw.New(c))
		return c.Image().(*image.RGBA)
	}
}
