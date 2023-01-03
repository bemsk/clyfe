package life

import (
	"fmt"
	"strings"
)

type Position interface {
	Lat() int
	Lon() int
	IsAdjecent(lat int, lon int) bool
	Vicinals() []Position
	String() string
}

type Cell struct {
	Position
	neighborCount int
}

type Life struct {
	population []*Cell
}

// New creates Life with zero population count
func New() *Life {
	p := []*Cell{}

	return  &Life{ population: p }
}

func (l *Life) String() string {
	s := []string{}

	for _, p := range l.population {
		s = append(s, fmt.Sprintf("%v", p))
	}
	
	return strings.Join(s, "\t")
}

func (l *Life) Tick() {
	seeds := l.seed()
	survivors := l.preserve()

	// reset all
	l.population = []*Cell{}

	for _, survivor := range survivors {
		l.Add(survivor.Position)
	}

	// create new cell
	for _, seed := range seeds {
		l.Add(seed)
	}
}

// Add creates new cell and add it into population
func (l *Life) Add(ps ...Position) {
	for _, p := range ps {
		var neighborCount int

		for _, existingCell := range l.population {
			if existingCell.IsAdjecent(p.Lat(), p.Lon()) {
				existingCell.neighborCount += 1
				neighborCount += 1
			}
		}

		c := &Cell{ p, neighborCount }
		
		l.population = append(l.population, c)
	}
}

func (l *Life) seed() []Position {
	s := []Position{}
	posCouter := make(map[string][]Position)

	for _, c := range l.population {
		neighborPositions := c.Vicinals()
		for _, np := range neighborPositions {
			posCouter[np.String()] = append(posCouter[np.String()], np)
		}
	}

	for _, pos := range posCouter {
		if len(pos) == 3 {
			s = append(s, pos[0])
		}
	}

	return s
}

func (l *Life) preserve() []*Cell {
	p := []*Cell{}

	for _, c := range l.population {
		if c.neighborCount > 1 && c.neighborCount < 4 {
			p = append(p, c)
		}
	}

	return p
}