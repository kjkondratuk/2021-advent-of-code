package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("inputs/day-4-small.txt")

	lines := strings.Split(string(bytes), "\n")

	turnLiteral := lines[0]
	turns := strings.Split(turnLiteral, ",")

	cards := make(map[int]Card)
	nextCard := 1
	for i := 1; i < len(lines); i++ {
		// skip lines which separate data
		if lines[i] == "" {
			continue
		}

		cards[nextCard-1] = NewCard()
		//fmt.Printf("Creating card\n")
		for x := 0; x < 5; x++ {
			ds := strings.ReplaceAll(strings.TrimPrefix(lines[i+x], " "), "  ", " ")
			numbers := strings.Split(ds, " ")
			if len(numbers) != 5 {
				panic("invalid card detected")
			}
			cards[nextCard-1].AppendRow(numbers)
		}
		i += 5
		nextCard++
	}

	fmt.Printf("Cards: %d\n", len(cards))

	firstWinner := -1
	var turnsTaken []string
	for i := 0; i < len(turns); i++ {
		turnsTaken = append(turnsTaken, turns[i])
		for ci, c := range cards {
			if filled := c.Mark(turns[i]); filled {
				if firstWinner < 0 {
					fmt.Printf("filled card: %d on turn %d\n", ci, i)
					fmt.Printf("turns: %+v\n", turnsTaken)
					fmt.Println("Card:")
					for _, row := range c.GetSpaces() {
						fmt.Printf("%+v\n", row)
					}
					firstWinner = ci
				}
			}
		}
	}
}

type card struct {
	spaces [][]string
	marks  [][]bool
}

type Card interface {
	AppendRow(s []string)
	Mark(s string) bool
	GetSpaces() [][]string
}

func NewCard() Card {
	return &card{
		spaces: make([][]string, 0),
		marks:  make([][]bool, 0),
	}
}

func (c *card) AppendRow(s []string) {
	c.spaces = append(c.spaces, s)

	marks := make([]bool, 0)
	for range s {
		marks = append(marks, false)
	}

	c.marks = append(c.marks, marks)
}

func (c *card) GetSpaces() [][]string {
	return c.spaces
}

func (c *card) Mark(s string) bool {
	for y, row := range c.spaces {
		for x, col := range row {
			if col == s {
				c.marks[y][x] = true

				completed := checkRow(c.marks[y])
				if completed {
					return completed
				} else {
					b := true
					for i := 0; i < 5; i++ {
						if c.marks[y][i] {
							b = false
						}
					}
					return b
				}
			}
		}
	}
	return false
}

func checkRow(s []bool) bool {
	for _, i := range s {
		if !i {
			return false
		}
	}
	return true
}
