package main

import (
	"fmt"
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

	//f := data.Filter(vertOrHorizontalFilter)
	f := data
	p := Plot{}
	for _, l := range f {
		p.Add(l)
	}
	p.Print(10, 10)

	overlapCount := 0
	for _, v := range p {
		if v >= 2 {
			overlapCount++
		}
	}

	fmt.Printf("Overlap Count: %d\n", overlapCount)
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

// Draw : draws a line on the provided plot
func (l Line) Draw(p *Plot) {
	//fmt.Printf("Drawing line from [%+v] to [%+v]\n", l.origin, l.dest)
	// fill in the origin on the plot
	if _, ok := (*p)[l.origin]; ok {
		(*p)[l.origin] += 1
	} else {
		(*p)[l.origin] = 1
	}
	vec := l.origin.RouteTo(l.dest)

	// iterate over the line in the direction of travel
	//fmt.Printf("Drawing vector: [%d] [%d]\n", vec.dist, vec.dir)
	intermediate := Point{}
	for i := 0; i < vec.dist; i++ {
		//fmt.Printf("Step: [%+v]\n", step)
		//fmt.Printf("Adding point: [%+v]\n", l.origin.Add(step))
		// fill line in the plot
		if i == 0 {
			intermediate = l.origin.Add(Point(vec.dir))
		} else {
			intermediate = intermediate.Add(Point(vec.dir))
		}
		if _, ok := (*p)[intermediate]; ok {
			(*p)[intermediate] += 1
		} else {
			(*p)[intermediate] = 1
		}
	}
}

var (
	Up        = Direction{0, -1}
	UpLeft    = Direction{-1, -1}
	UpRight   = Direction{1, -1}
	Down      = Direction{0, 1}
	DownLeft  = Direction{-1, 1}
	DownRight = Direction{1, 1}
	Right     = Direction{1, 0}
	Left      = Direction{-1, 0}
	None      = Direction{0, 0}
)

type Direction Point

type Point struct {
	x int
	y int
}

// Add : adds the values of one point to another to provide a new point
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

// DirTo : determines the direction of the destination point from the source point
func (p Point) DirTo(d Point) Direction {
	switch {
	case p.y > d.y && p.x < d.x:
		return UpRight
	case p.y > d.y && p.x > d.x:
		return UpLeft
	case p.y < d.y && p.x > d.x:
		return DownLeft
	case p.y < d.y && p.x < d.x:
		return DownRight
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

// RouteTo : determines a Vector from the source point to the destination
func (p Point) RouteTo(d Point) Vector {
	var val int
	dir := p.DirTo(d)
	switch dir {
	case UpRight:
		val = p.y - d.y
	case UpLeft:
		val = p.y - d.y
	case Up:
		val = p.y - d.y
	case DownRight:
		val = d.y - p.y
	case DownLeft:
		val = d.y - p.y
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

// Add : adds a line to the plot (by drawing it)
func (p *Plot) Add(l Line) {
	l.Draw(p)
}

// Print : write the plot to standard output at the given size
func (p *Plot) Print(w int, h int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if val, ok := (*p)[Point{x: j, y: i}]; ok {
				fmt.Printf(lib.Red("%d"), val)
			} else {
				fmt.Printf(".")
			}

			if (j+1)%w == 0 {
				fmt.Printf("\n")
			}
		}
	}
}
