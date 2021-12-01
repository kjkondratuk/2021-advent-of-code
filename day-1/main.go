package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("inputs/day-1.txt")
	if err != nil {
		log.Fatalf("Error reading input file: [%s]", err)
	}

	lines := strings.Split(string(f), "\n")
	incDepths := Run(lines)

	log.Printf("Entries > previous: %d\n", incDepths)
}

func Run(lines []string) uint32 {
	var acc uint32 = 0
	var prev *uint32
	for _, l := range lines {
		d, err := strconv.ParseUint(l, 10, 32)
		depth := uint32(d)
		if err != nil {
			log.Fatalf("Could not parse input to uint32: [%s]", l)
		}

		if prev != nil {
			if depth > *prev {
				log.Printf("[%d] > [%d]", depth, *prev)
				acc++
			} else {
				log.Printf("* [%d] <= [%d] *", depth, *prev)
			}
		} else {
			var zero uint32 = 0
			prev = &zero
			log.Printf("First record...")
		}
		prev = &depth
	}
	return acc
}
