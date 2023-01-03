package main

import (
	"fmt"

	"github.com/bemsk/clyfe/cartesian"
	"github.com/bemsk/clyfe/life"
)

func display(s fmt.Stringer) {
	fmt.Printf("population:\n")
	fmt.Printf("%v", s)
	fmt.Printf("\n")
}

func main() {
	coordinates := [][]int{
		{0, -1},
		{0, 0},
		{0, 1 },
	}

	points := []life.Position{}

	for _, c := range coordinates {
		p := cartesian.NewPoint(c[0], c[1])
		points = append(points, p)
	}

	clyfe := life.New()
	clyfe.Add(points...)

	for i := 0; i < 10; i ++ {
		clyfe.Tick()
		display(clyfe)
	}
}