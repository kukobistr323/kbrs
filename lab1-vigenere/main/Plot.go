package main

import "github.com/Arafatk/glot"

const Dimension = 2

func drawPlot(points [][]float64, xLabel string, yLabel string, maxX int, maxY int, path string) {
	plot, _ := glot.NewPlot(Dimension, false, false)
	plot.SetTitle("Results")
	plot.SetXLabel(xLabel)
	plot.SetYLabel(yLabel)
	plot.SetXrange(0, maxX)
	plot.SetYrange(0, maxY)
	plot.AddPointGroup("Points", "point", points)
	plot.SavePlot(path)
}

func preparePointsFirstPlot(result []int) [][]float64 {
	points := make([][]float64, Dimension)

	for i := 0; i < Dimension; i++ {
		points[i] = make([]float64, Tests, Tests)
	}

	temp := 0
	for i := 0; i < Tests; i++ {
		for j := 0; j < Tests; j++ {
			temp += result[i*10+j]
		}
		points[0][i] = float64(len(Keys[i]))
		points[1][i] = float64(temp) / float64(Tests*Tests)
		temp = 0
	}
	return points
}

func preparePointsSecondPlot(result []int) [][]float64 {
	points := make([][]float64, Dimension)

	for i := 0; i < Dimension; i++ {
		points[i] = make([]float64, Tests, Tests)
	}

	temp := 0
	for i := 0; i < Tests; i++ {
		for j := 0; j < Tests; j++ {
			temp += result[i+j*10]
		}
		points[0][i] = float64((i + 1) * 1000)
		points[1][i] = float64(temp) / float64(Tests*Tests)
		temp = 0
	}
	return points
}
