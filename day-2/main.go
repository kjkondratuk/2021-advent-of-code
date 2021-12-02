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
	navigator                = NewNavigatorWithPosition(Postion{}, map[string]NavigatorFunc{
		"up": func(scalar int, p Postion) Postion {
			p.y -= scalar
			return p
		}, "down": func(scalar int, p Postion) Postion {
			p.y += scalar
			return p
		}, "forward": func(scalar int, p Postion) Postion {
			p.x += scalar
			return p
		},
	})
	aimNavigator = NewNavigatorWithPosition(Postion{}, map[string]NavigatorFunc{
		"up": func(scalar int, p Postion) Postion {
			p.a -= scalar
			return p
		}, "down": func(scalar int, p Postion) Postion {
			p.a += scalar
			return p
		}, "forward": func(scalar int, p Postion) Postion {
			p.x += scalar
			p.y += p.a * scalar
			return p
		},
	})
)

type Postion struct {
	x int
	y int
	a int
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

type NavigatorFunc func(int, Postion) Postion

type Navigator struct {
	pos Postion
	ops map[string]NavigatorFunc
}

func NewNavigatorWithPosition(start Postion, ops map[string]NavigatorFunc) Navigator {
	return Navigator{
		pos: start,
		ops: ops,
	}
}

func (n *Navigator) Go(ins Instruction) Postion {
	//log.Printf("going: %s %d", ins.direction, ins.scalar)
	n.pos = n.ops[ins.direction](ins.scalar, n.pos)
	return n.pos
}

func main() {
	data := lib.ReadData("inputs/day-2.txt")
	var pos Postion
	var aimPos Postion
	for _, dir := range data {
		i, err := InstructionFromString(dir)
		if err != nil {
			panic(err)
		}

		pos = navigator.Go(i)
		aimPos = aimNavigator.Go(i)
		//log.Printf("New Postion: %+v", pos)
	}

	log.Printf("Final Navigator Location: %+v", pos)
	log.Printf("Final Aim Navigator Location: %+v", aimPos)
}
