package main

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/kjkondratuk/2021-advent-of-code/lib"
)

var (
	ErrInvalidPositionString = errors.New("invalid position string")
	navigator                = NewNavigatorWithPosition(Position{
		x: 0,
		y: 0,
	}, map[string]NavigatorFunc{
		"up": func(scalar int, p Position) Position {
			p.y -= scalar
			return p
		}, "down": func(scalar int, p Position) Position {
			p.y += scalar
			return p
		}, "forward": func(scalar int, p Position) Position {
			p.x += scalar
			return p
		},
	})
)

type Position struct {
	x int
	y int
}

type Instruction struct {
	direction string
	scalar    int
}

func InstructionFromString(s string) (Instruction, error) {
	if s == "" {
		return Instruction{}, ErrInvalidPositionString
	}
	vector := strings.Split(s, " ")
	if len(vector) != 2 {
		return Instruction{}, ErrInvalidPositionString
	} else {
		dist, err := strconv.ParseInt(vector[1], 10, 64)
		if err != nil {
			return Instruction{}, err
		}
		d := int(dist)
		return Instruction{
			direction: vector[0],
			scalar:    d,
		}, nil
	}
}

type NavigatorFunc func(int, Position) Position

type Navigator struct {
	pos Position
	ops map[string]NavigatorFunc
}

func NewNavigatorWithPosition(start Position, ops map[string]NavigatorFunc) Navigator {
	return Navigator{
		pos: start,
		ops: ops,
	}
}

func (n *Navigator) Go(ins Instruction) Position {
	//log.Printf("going: %s %d", ins.direction, ins.scalar)
	n.pos = n.ops[ins.direction](ins.scalar, n.pos)
	return n.pos
}

func main() {
	data := lib.ReadData("inputs/day-2.txt")
	var pos Position
	for _, dir := range data {
		i, err := InstructionFromString(dir)
		if err != nil {
			panic(err)
		}

		pos = navigator.Go(i)
		//log.Printf("New Position: %+v", pos)
	}

	log.Printf("Final Location: %+v", pos)
}
