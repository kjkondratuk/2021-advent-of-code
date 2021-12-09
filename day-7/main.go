package main

import (
	"fmt"
	"github.com/kjkondratuk/2021-advent-of-code/lib"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := lib.NewDataReader("inputs/day-7.txt").Read()
	lines := data.([]string)
	posStr := strings.Split(lines[0], ",")
	pos := make(IntCollection, len(posStr))
	for i, s := range posStr {
		p, _ := strconv.Atoi(s)
		pos[i] = p
	}

	fmt.Printf("Numbers: %+v\n", pos)
	sort.Ints(pos)
	fmt.Printf("Numbers: %+v\n", pos)

	max := pos[len(pos)-1:]
	costs := make(IntCollection, max[0])
	factCosts := make([]int, max[0])

	for i := 0; i < max[0]; i++ {
		cost := pos.CostTo(i)
		factCost := pos.GradientCostTo(i)
		costs[i] = cost
		factCosts[i] = factCost
	}

	for i, c := range costs {
		fmt.Printf("Cost %d = %d\n", i, c)
		fmt.Printf("Fact Cost %d = %d\n", i, factCosts[i])
	}
	sort.Ints(costs)
	sort.Slice(factCosts, func(i, j int) bool {return factCosts[i] < factCosts[j]})
	fmt.Printf("Least expensive: %d\n", costs[0])
	fmt.Printf("Least expensive %d\n", factCosts[0])

}

type IntCollection []int

func (l IntCollection) CostTo(n int) int {
	acc := 0
	for _, x := range l {
		if x > n {
			acc += x - n
		} else if x < n {
			acc += n - x
		}
	}
	return acc
}

func (l IntCollection) GradientCostTo(n int) int {
	acc := 0
	for _, p := range l {
		if p > n {
			acc += (p - n) * ((p - n) + 1) / 2
		} else if p < n {
			acc += (n - p) * ((n - p) + 1) / 2
		}
	}

	return acc
}
