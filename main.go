package main

import (
	"fmt"
	"math"
)

type row struct {
	x float64
	y float64
}

type table []row

func (t *table) means() (float64, float64) {
	xMean := 0.0
	yMean := 0.0
	for _, row := range *t {
		xMean += row.x
		yMean += row.y
	}
	return xMean / float64(len(*t)), yMean / float64(len(*t))
}

func (t *table) ols() (float64, float64) {
	xMean, yMean := t.means()
	num, den := 0.0, 0.0
	for _, row := range *t {
		num += (row.x - xMean) * (row.y - yMean)
		den += (math.Pow((row.x - xMean), 2))
	}
	m := num / den
	b := yMean - m*xMean
	return m, b
}

func main() {
	t := table{
		{99, 54.02434},
		{2, 3.32323},
	}

	m, b := t.ols()
	fmt.Printf("y = %.3fx + %.3f\n", m, b)

}
