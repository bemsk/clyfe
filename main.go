package main

import (
	"github.com/bemsk/clyfe/cartesian"
	"github.com/bemsk/clyfe/life"
)

func main() {
	coordinates := [][]int{
		{0, -1},
		{0, 0},
		{0, 1 },
	}

	points := []*cartesian.Point{}

	for _, c := range coordinates {
		p := cartesian.NewPoint(c[0], c[1])
		points = append(points, p)
	}

	clyfe := life.New()

	// passing []*cartesian.Point apparently doesn't work
	// see https://www.timr.co/go-interfaces-the-tricky-parts/
	clyfe.Add(points...)

	clyfe.Tick()
}