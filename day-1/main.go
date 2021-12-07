package main

import (
	"log"
	"strconv"

	"github.com/kjkondratuk/2021-advent-of-code/lib"
)

func main() {
	lines := lib.NewDataReader("inputs/day-1.txt").Read().([]string)
	var p lib.IntCollection = parse(lines)
	log.Printf("Parsed Length: [%d]", len(p))
	c := p.Comprehend(3, func(set []int) []int {
		//log.Printf("Summarizing: %+v", set)
		var acc = 0
		for i := 0; i < len(set); i++ {
			acc += set[i]
		}
		//log.Printf("Result: %d", acc)
		return []int{acc}
	})
	log.Printf("Comprehension Length: [%d]", len(c))
	incDepths := determineIncreasingPeriods(c)

	log.Printf("Entries > previous: %d\n", incDepths)
}

func parse(lines []string) []int {
	res := make([]int, 0)
	for _, l := range lines {
		d, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("Could not parse input to uint32: [%s]", l)
		}
		res = append(res, d)
	}
	return res
}

func determineIncreasingPeriods(depths []int) int {
	var acc = 0
	var prev *int
	for _, d := range depths {
		if prev != nil {
			if d > *prev {
				acc++
			}
		} else {
			var zero = 0
			prev = &zero
			log.Printf("First record...")
		}
		*prev = d
	}
	return acc
}
