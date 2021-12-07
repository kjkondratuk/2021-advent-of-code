package main

import (
	"fmt"
	"github.com/kjkondratuk/2021-advent-of-code/lib"
	"strconv"
	"strings"
)

var (
	loader = func(lines []string) interface{} {
		turnLiteral := lines[0]
		turns := strings.Split(turnLiteral, ",")

		cards := make([]Card, 0)
		nextCard := 1
		for j := 1; j < len(lines); j++ {
			// skip lines which separate data
			if lines[j] == "" {
				continue
			}

			cards = append(cards, NewCard(nextCard-1))
			//fmt.Printf("Creating card\n")
			for x := 0; x < 5; x++ {
				ds := strings.ReplaceAll(strings.TrimPrefix(lines[j+x], " "), "  ", " ")
				numbers := strings.Split(ds, " ")
				if len(numbers) != 5 {
					panic("invalid card detected")
				}
				cards[nextCard-1].AppendRow(numbers)
			}
			j += 5
			nextCard++
		}

		result := make([]interface{}, 0)
		result = append(result, turns, cards)

		return result
	}
)

func main() {
	//turns, cards := parseFile("inputs/day-4.txt")
	data := lib.NewDataReaderWithLoader("inputs/day-4.txt", loader).Read()
	params := data.([]interface{})
	turns := params[0].([]string)
	cards := params[1].([]Card)

	fmt.Printf("Cards: %d\n", len(cards))
	//fmt.Printf("Turns: %+v\n", turns)

	turnIter := NewTurnIterator(turns)

	var winner Card
	for turnIter.HasNext() && winner == nil {
		winner = turnIter.DoTurn(cards)
		//fmt.Printf("Doing turn: %d\n", turnIter.GetIndex())
	}
	if winner != nil {
		winner.Print()
		fmt.Printf("Total: %d\n", winner.TotalUnmarked())
		fmt.Printf("Last Filled: %d\n", turnIter.Current())
		fmt.Printf("Score: %d\n\n", winner.TotalUnmarked()*turnIter.Current())
	}

	// parsing again because cards are mutable and we've dirtied the previous set
	d := lib.NewDataReaderWithLoader("inputs/day-4.txt", loader).Read()
	p := d.([]interface{})
	t := p[0].([]string)
	c := p[1].([]Card)

	loserIter := NewTurnIterator(t)

	var winners []Card
	for loserIter.HasNext() && len(winners) < len(c) {
		w := loserIter.DoTurn(c)
		if w != nil && !contains(winners, w) {
			winners = append(winners, w)
		}
		//fmt.Printf("Doing turn: %d\n", loserIter.GetIndex())
	}
	if len(winners) > 0 {
		fmt.Printf("Winners: %d\n", len(winners))
		fmt.Printf("Last Winner:\n")
		lastWinner := winners[len(winners)-1:][0]
		lastWinner.Print()
		fmt.Printf("Total: %d\n", lastWinner.TotalUnmarked())
		fmt.Printf("Last Filled: %d\n", loserIter.Current())
		fmt.Printf("Score: %d\n\n", lastWinner.TotalUnmarked()*loserIter.Current())
	}
}

func contains(cards []Card, card Card) bool {
	for _, c := range cards {
		if c.Name() == card.Name() {
			return true
		}
	}
	return false
}

type card struct {
	name     int
	spaces   [][]string
	marks    [][]bool
	complete bool
}

type Card interface {
	AppendRow(s []string)
	Mark(s string) bool
	Print()
	Complete() bool
	TotalUnmarked() int
	Name() int
}

func NewCard(name int) Card {
	return &card{
		name:   name,
		spaces: make([][]string, 0),
		marks:  make([][]bool, 0),
	}
}

func (c *card) Name() int {
	return c.name
}

func (c *card) Complete() bool {
	return c.complete
}

func (c *card) AppendRow(s []string) {
	c.spaces = append(c.spaces, s)

	marks := make([]bool, 0)
	for range s {
		marks = append(marks, false)
	}

	c.marks = append(c.marks, marks)
}

// Mark : marks a position on a card if available.  returns true if this mark completed a row/column
func (c *card) Mark(s string) bool {
	for ri, row := range c.spaces {
		for ci, item := range row {
			if s == item {
				c.marks[ri][ci] = true
				//fmt.Printf("Marking [%d] [%d] on card [%d]\n", ri, ci, c.name)
				complete := c.RowCompleted(ri) || c.ColCompleted(ci)
				if complete {
					c.complete = true
				}
				return complete
			}
		}
	}
	return false
}

func (c *card) TotalUnmarked() int {
	total := 0
	for ri, row := range c.spaces {
		for ci, item := range row {
			if !c.marks[ri][ci] {
				val, _ := strconv.Atoi(item)
				total += val
			}
		}
	}
	return total
}

func (c *card) RowCompleted(row int) bool {
	for _, i := range c.marks[row] {
		if i != true {
			return false
		}
	}
	return true
}

func (c *card) ColCompleted(col int) bool {
	for _, row := range c.marks {
		if row[col] != true {
			return false
		}
	}
	return true
}

func (c *card) Print() {
	fmt.Printf("Card: %d\n", c.name)
	for ri, row := range c.spaces {
		for ci, item := range row {
			str := item
			if len(str) == 1 {
				str = " " + str
			}
			if c.marks[ri][ci] {
				fmt.Printf(lib.Red(str) + " ")
			} else {
				fmt.Printf(str + " ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

type TurnIterator struct {
	turns   []string
	current int
}

func NewTurnIterator(turns []string) TurnIterator {
	return TurnIterator{
		turns:   turns,
		current: 0,
	}
}

// DoTurn : performs a turn across cards, in order and increments the array pointer.
// Returns true if a card was completed, false otherwise.
func (i *TurnIterator) DoTurn(cards []Card) Card {
	for j := 0; j < len(cards); j++ {
		if !cards[j].Complete() {
			if completed := cards[j].Mark(i.turns[i.current]); completed {
				//cards[j].Print()
				//i.current++
				return cards[j]
			}
		}
	}
	i.current++
	return nil
}

func (i *TurnIterator) GetIndex() int {
	return i.current
}

func (i *TurnIterator) Current() int {
	v, _ := strconv.Atoi(i.turns[i.current])
	return v
}

func (i *TurnIterator) HasNext() bool {
	return i.current < len(i.turns)-1
}
