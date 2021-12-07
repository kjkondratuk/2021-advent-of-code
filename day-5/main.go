package main

import (
	"github.com/kjkondratuk/2021-advent-of-code/lib"
	"strconv"
	"strings"
)

var (
	// loads the data based on lines and the predetermined format
	loader = func(lines []string) interface{} {
		l := make([]Line, 0)
		for _, line := range lines {
			points := strings.Split(line, " -> ")
			one := strings.Split(points[0], ",")
			two := strings.Split(points[1], ",")
			x1s, _ := strconv.Atoi(one[0])
			y1s, _ := strconv.Atoi(one[1])
			x2s, _ := strconv.Atoi(two[0])
			y2s, _ := strconv.Atoi(two[1])
			l = append(l, Line{
				origin: Point{x1s, y1s},
				dest:   Point{x2s, y2s},
			})
		}
		return l
	}

	// filters a LineList for lines that are either vertical or horizontal
	vertOrHorizontalFilter = func(list LineList) LineList {
		res := make(LineList, 0)
		for _, i := range list {
			if i.origin.y == i.dest.y || i.origin.x == i.dest.x {
				res = append(res, i)
			}
		}
		return res
	}
)

func main() {
	d := lib.NewDataReaderWithLoader("inputs/day-5.txt", loader).Read()
	data := LineList(d.([]Line))

	f := data.Filter(vertOrHorizontalFilter)
	p := Plot{}
	for _, l := range f {
		p.Add(l)
	}
	p.Print()
}

type LineList []Line

type LinePredicate func(LineList) LineList

func (l LineList) Filter(predicate LinePredicate) []Line {
	return predicate(l)
}

type Line struct {
	origin Point
	dest   Point
}

func (l Line) Draw(p *Plot) {
	// fill in the origin on the plot
	if _, ok := (*p)[l.origin]; ok {
		(*p)[l.origin] += 1
	} else {
		(*p)[l.origin] = 1
	}
	vec := l.origin.RouteTo(l.dest)
	step := Point{
		x: 0, y: 0,
	}
	switch vec.dir {
	case Up:
		step.x--
	case Down:
		step.x++
	case Left:
		step.y--
	case Right:
		step.y++
	case None:
	}

	// iterate over the line in the direction of travel
	for i := 0; i < vec.dist; i++ {
		// fill line in the plot
		if _, ok := (*p)[l.origin.Add(step)]; ok {
			(*p)[l.origin.Add(step)] += 1
		} else {
			(*p)[l.origin.Add(step)] = 1
		}
	}
}

const (
	Up = iota
	Down
	Right
	Left
	None
)

type Direction int

type Point struct {
	x int
	y int
}

func (p Point) Add(d Point) Point {
	cp := p
	cp.y += d.y
	cp.x += d.x
	return cp
}

type Vector struct {
	dist int
	dir  Direction
}

func (p Point) DirTo(d Point) Direction {
	switch {
	case p.x > d.x:
		return Left
	case p.y > d.y:
		return Up
	case p.x < d.x:
		return Right
	case p.y < d.y:
		return Down
	default:
		return None
	}
}

func (p Point) RouteTo(d Point) Vector {
	var val int
	dir := p.DirTo(d)
	switch dir {
	case Up:
		val = p.y - d.y
	case Down:
		val = d.y - p.y
	case Left:
		val = p.x - d.x
	case Right:
		val = d.x - p.x
	case None:
		val = 0
	}
	return Vector{
		dist: val,
		dir:  dir,
	}
}

type Plot map[Point]int

func (p *Plot) Add(l Line) {
	l.Draw(p)
}

func (p *Plot) Print() {

}
