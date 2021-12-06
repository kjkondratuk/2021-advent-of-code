package main

import (
	"fmt"
	"strconv"
)

var (
	completedCardCount = 0
)

func main() {
	turns, cards := parseFile("inputs/day-4.txt")
	fmt.Printf("Cards: %d\n", len(cards))
	fmt.Printf("Turns: %+v\n", turns)

	s := score(turns, cards, true)
	fmt.Printf("Score 1: %d\n\n", s)

	s1 := score(turns, cards, false)
	fmt.Printf("Score 2: %d\n\n", s1)
}

func score(turns []string, cards map[int]Card, firstWinner bool) int {
	turn, winner := determineWinner(turns, cards, firstWinner)

	for i, c := range cards {
		fmt.Printf("\nIndex: %d - Card:\n", i)
		for _, row := range c.GetSpaces() {
			fmt.Println(row)
		}
	}
	fmt.Printf("\nWinner: %d - Card:\n", winner)
	for _, row := range cards[winner].GetSpaces() {
		fmt.Println(row)
	}

	fmt.Printf("\nTurn won: %d\n", turn)
	winNum, _ := strconv.Atoi(turns[turn-1])
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
				if first && winner < 0 {
					winner = ci
					goto winnerEnd
				} else if !first {
					if completedCardCount >= len(cards) {
						winner = ci
						goto winnerEnd
					}
				}
			}
		}
	}

winnerEnd:
	return len(turnsTaken), winner
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

func checkRow(s []bool) bool {
	for _, i := range s {
		if !i {
			return false
		}
	}
	return true
}

func checkBingo(c *card, ri int, ci int) bool {
	if rowBingo := checkRow(c.marks[ri]); rowBingo {
		fmt.Printf("rowBingo based on row %s \n", c.spaces[0][0])
		completedCardCount++
		fmt.Printf("completed card count: %d\n", completedCardCount)
		return rowBingo
	} else {
		// check for column bingo
		colBingo := true
		fmt.Printf("Checking column bingo:\n")
		for i := 0; i < 5; i++ {
			if !c.marks[i][ci] {
				colBingo = false
				fmt.Printf("No ColBingo - row: %d col: %d - %t\n", ri, i, c.marks[ri][i])
				break
			} else {
				fmt.Printf("ColBingo - row: %d col: %d - %t\n", ri, i, c.marks[ri][i])
			}
		}
		if colBingo {
			fmt.Printf("rowBingo based on column %s \n", c.spaces[0][0])
			completedCardCount++
			fmt.Printf("completed card count: %d\n", completedCardCount)
		}
		return colBingo
	}
}
