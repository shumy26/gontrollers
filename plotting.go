package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func plotting(ch chan float64) {
	plt := plot.New()
	plt.Title.Text = "PID controller test"
	plt.X.Label.Text = "Time"
	plt.Y.Label.Text = "Controlled Variable"

	points := make([]float64, 0)

	for {
		value, ok := <-ch
		if !ok {
			break
		}
		points = append(points, value)
	}

	pltter := make(plotter.XYs, len(points))

	for i := range len(points) {
		pltter[i].X = float64(i)
		pltter[i].Y = points[i]
	}

	line, err := plotter.NewLine(pltter)
	if err != nil {
		log.Fatal("err")
	}
	plt.Add(line)

	plt.Save(6*vg.Inch, 4*vg.Inch, "pid_plot.png")
}
