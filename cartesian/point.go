package cartesian

import (
	"math"
	"strconv"

	"github.com/bemsk/clyfe/life"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) life.Position {
	return &Point{ x, y }
}

func (p *Point) Lat() int {
	return p.X
}

func (p *Point) Lon() int {
	return p.Y
}

func (p *Point) IsAdjecent(x int, y int) bool {
	distX := math.Abs(float64(p.X - x))
	distY := math.Abs(float64(p.Y - y))

	return distX < 2 && distY < 2
}

func (p *Point) Vicinals() []life.Position {
	v := []life.Position{}

	directions := [8][2]int{
		{0, 1}, 	// N
		{1, 1}, 	// NE
		{1, 0},		// E
		{1, -1},	// SE
		{0, -1},	// S
		{-1, -1},	// SW
		{-1, 0},	// W
		{-1, 1},	// NW
	}

	for _, d := range directions {
		v = append(v, NewPoint(p.X - d[0], p.Y - d[1]))
	}

	return v
}

func (p *Point) String() string {
	strX := strconv.Itoa(p.X)
	strY := strconv.Itoa(p.Y)

	return strX + "," + strY
}