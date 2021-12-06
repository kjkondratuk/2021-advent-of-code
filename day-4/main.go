package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	turns, cards := parseFile("inputs/day-4-small.txt")
	fmt.Printf("Cards: %d\n", len(cards))
	fmt.Printf("Turns: %+v\n", turns)

	s := score(turns, cards, true)
	fmt.Printf("Score 1: %d\n\n", s)

	s1 := score(turns, cards, false)
	fmt.Printf("Score 2: %d\n\n", s1)
}

func score(turns []string, cards map[int]Card, firstWinner bool) int {
	turn, winner := determineWinner(turns, cards, firstWinner)
	fmt.Printf("\nTurn won: %d\n", turn)
	fmt.Printf("Winner index: %d\n", winner)

	fmt.Printf("Index: %d - Card: %+v\n", winner, cards[winner].GetSpaces())
	winNum, _ := strconv.Atoi(turns[turn - 1])
	s := cards[winner].Score(winNum)
	return s
}

func determineWinner(turns []string, cards map[int]Card, first bool) (int, int) {
	winner := -1
	var turnsTaken []string
	for i := 0; i < len(turns); i++ {
		turnsTaken = append(turnsTaken, turns[i])
		for ci, c := range cards {
			if filled := c.Mark(turns[i]); filled {
				fmt.Printf("completions: %d\n", cards[ci].GetCompletions())
				if !first && cards[ci].GetCompletions() == 1 {
					winner = ci
					fmt.Printf("updating winner: %d\n", winner)
				}
				if first && winner < 0 {
					winner = ci
					goto winnerEnd
				}
			}
		}
	}

	winnerEnd:
	return len(turnsTaken), winner
}

func parseFile(file string) ([]string, map[int]Card) {
	bytes, _ := ioutil.ReadFile(file)

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

	return turns, cards
}

type card struct {
	spaces [][]string
	marks  [][]bool
	completeCount int
}

type Card interface {
	AppendRow(s []string)
	Mark(s string) bool
	GetSpaces() [][]string
	Score(turn int) int
	Complete()
	GetCompletions() int
}

func NewCard() Card {
	return &card{
		spaces: make([][]string, 0),
		marks:  make([][]bool, 0),
		completeCount: 0,
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
	colorizedSpaces := c.spaces

	for ri, row := range colorizedSpaces {
		for ci, v := range row {
			if c.marks[ri][ci] {
				colorizedSpaces[ri][ci] = "\033[31m" + v + "\033[0m"
			}
		}
	}

	return colorizedSpaces
}

func (c *card) Mark(s string) bool {
	for ri, row := range c.spaces {
		for ci, col := range row {
			if col == s {
				// TODO : started adding completions to try to stop when we've reached the final first completion
				if c.completeCount == 0 {
					c.marks[ri][ci] = true
				}
				completed := checkRow(c.marks[ri])
				if completed {
					//fmt.Printf("completed based on row\n")
					c.Complete()
					return completed
				} else {
					b := true
					for i := 0; i < 5; i++ {
						if c.marks[ri][i] {
							b = false
						}
					}
					if b {
						//fmt.Printf("columns checked\n")
						c.Complete()
					}
					return b
				}
			}
		}
	}
	return false
}

func (c *card) Score(winnerNum int) int {
	total := 0
	for ri, row := range c.marks {
		for ci, marked := range row {
			if !marked {
				num, _ := strconv.Atoi(c.spaces[ri][ci])
				//fmt.Printf("adding: %d\n", num)
				total += num
			} else {
				//fmt.Printf("not marked\n")
			}
		}
	}
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Turn: %d\n", winnerNum)
	return total * winnerNum
}

func (c *card) Complete() {
	c.completeCount++
}

func (c *card) GetCompletions() int {
	return c.completeCount
}

func checkRow(s []bool) bool {
	for _, i := range s {
		if !i {
			return false
		}
	}
	return true
}
