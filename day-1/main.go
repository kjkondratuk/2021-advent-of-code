package main

import (
	"log"
	"strconv"

	"github.com/kjkondratuk/2021-advent-of-code/lib"
)

func main() {
	lines := lib.ReadData("inputs/day-1.txt")
	var p lib.UInt32Comprehendable = parse(lines)
	log.Printf("Parsed Length: [%d]", len(p))
	c := p.Comprehend(3, func(set []uint32) []uint32 {
		//log.Printf("Summarizing: %+v", set)
		var acc uint32 = 0
		for i := 0; i < len(set); i++ {
			acc += set[i]
		}
		//log.Printf("Result: %d", acc)
		return []uint32{acc}
	})
	log.Printf("Comprehension Length: [%d]", len(c))
	incDepths := determineIncreasingPeriods(c)

	log.Printf("Entries > previous: %d\n", incDepths)
}

func parse(lines []string) []uint32 {
	res := make([]uint32, 0)
	for _, l := range lines {
		d, err := strconv.ParseUint(l, 10, 32)
		depth := uint32(d)
		if err != nil {
			log.Fatalf("Could not parse input to uint32: [%s]", l)
		}
		res = append(res, depth)
	}
	return res
}

func determineIncreasingPeriods(depths []uint32) uint32 {
	var acc uint32 = 0
	var prev *uint32
	for _, d := range depths {
		//log.Printf("evaluating: [%d]", d)
		if prev != nil {
			if d > *prev {
				//log.Printf("[%d] > [%d]", d, *prev)
				acc++
			} /* else {
				log.Printf("* [%d] <= [%d] *", d, *prev)
			}*/
		} else {
			var zero uint32 = 0
			prev = &zero
			log.Printf("First record...")
		}
		*prev = d
	}
	return acc
}
