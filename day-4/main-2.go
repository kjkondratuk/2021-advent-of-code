package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	turns, cards := parseFile("inputs/day-4.txt")
	fmt.Printf("Cards: %d\n", len(cards))
	fmt.Printf("Turns: %+v\n", turns)

	turnIter := NewTurnIterator(turns)

	var winner Card
	for turnIter.HasNext() && winner == nil {
		winner = turnIter.DoTurn(cards)
		fmt.Printf("Doing turn: %d\n", turnIter.GetIndex())
	}
	if winner != nil {
		fmt.Printf("Total: %d\n", winner.TotalUnmarked())
		fmt.Printf("Last Filled: %d\n", turnIter.Current())
		fmt.Printf("Score: %d\n\n", winner.TotalUnmarked()*turnIter.Current())
	}

	t, c := parseFile("inputs/day-4.txt")
	loserIter := NewTurnIterator(t)

	var winners []Card
	for loserIter.HasNext() && len(winners) < len(c) {
		w := loserIter.DoTurn(c)
		if w != nil && !contains(winners, w) {
			winners = append(winners, w)
		}
		fmt.Printf("Doing turn: %d\n", loserIter.GetIndex())
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

func parseFile(file string) ([]string, []Card) {
	bytes, _ := ioutil.ReadFile(file)

	lines := strings.Split(string(bytes), "\n")

	turnLiteral := lines[0]
	turns := strings.Split(turnLiteral, ",")

	cards := make([]Card, 0)
	nextCard := 1
	for i := 1; i < len(lines); i++ {
		// skip lines which separate data
		if lines[i] == "" {
			continue
		}

		cards = append(cards, NewCard(nextCard-1))
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

	return turns, cards
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
				fmt.Printf("\033[31m" + str + "\033[0m" + " ")
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
		} else {
			fmt.Printf("Card %d complete, skipping marking\n", cards[j].Name())
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
	hasNext := i.current < len(i.turns)-1
	fmt.Printf("Checking hasNext: %t - %d %d\n", hasNext, i.current, len(i.turns)-1)
	return hasNext
}
