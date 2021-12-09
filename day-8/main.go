package main

import (
	"fmt"
	"strings"

	"github.com/kjkondratuk/2021-advent-of-code/day-8/string_sort"

	"github.com/kjkondratuk/2021-advent-of-code/day-8/translate"
	"github.com/kjkondratuk/2021-advent-of-code/lib"
)

var (
	loader = func(lines []string) interface{} {
		problems := make([]Problem, len(lines))
		for i, line := range lines {
			parts := strings.Split(line, " | ")
			patternChunk := strings.ReplaceAll(parts[0], "\r", "")
			outValChunk := strings.ReplaceAll(parts[1], "\r", "")
			patterns := strings.Split(patternChunk, " ")
			for j, p := range patterns {
				patterns[j] = string_sort.String(p)
			}
			outVals := strings.Split(outValChunk, " ")
			for j, o := range patterns {
				patterns[j] = string_sort.String(o)
			}
			problems[i] = Problem{
				Patterns: patterns,
				OutVals:  outVals,
			}
		}
		return problems
	}

	// BEHOLD!  My laziness knows no bounds!
	stringToNumber = translate.ToInt{
		string_sort.String("abcefg"):  0,
		string_sort.String("cf"):      1,
		string_sort.String("acdeg"):   2,
		string_sort.String("acdfg"):   3,
		string_sort.String("bcdf"):    4,
		string_sort.String("abdfg"):   5,
		string_sort.String("abdefg"):  6,
		string_sort.String("acf"):     7,
		string_sort.String("abcdefg"): 8,
		string_sort.String("abcdfg"):  9,
	}

	// Part 2 : the re-baddening!
	numberToString = translate.ToString{
		0: string_sort.String("abcefg"),
		1: string_sort.String("cf"),
		2: string_sort.String("acdeg"),
		3: string_sort.String("acdfg"),
		4: string_sort.String("bcdf"),
		5: string_sort.String("abdfg"),
		6: string_sort.String("abdefg"),
		7: string_sort.String("acf"),
		8: string_sort.String("abcdefg"),
		9: string_sort.String("abcdfg"),
	}

	baseUniqueCounts = []int{
		1, 4, 7, 8,
	}
)

func main() {
	data := lib.NewDataReaderWithLoader("inputs/day-8-small.txt", loader).Read()
	problems := data.([]Problem)

	overallCount := 0
	for i, problem := range problems {
		count := 0
		// TODO : need to add unique count items to the base set here before we check the output
		// values for unique values

		//uniqueCounts := baseUniqueCounts
		//for _, p := range problem.Patterns {
		//
		//}
		for _, ns := range problem.OutVals {
			fmt.Printf("%s", ns)

			//if len(ns) == 1 || len(ns) == 4 || len(ns) == 7 || len(ns) == 8 {
			//	fmt.Printf(" Incrementing")
			//	count++
			//}
			fmt.Printf("\n")
		}
		overallCount += count
		fmt.Printf("Problem %d: %d\n", i, count)
	}
	fmt.Printf("Total: %d", overallCount)
}

type Problem struct {
	Patterns []string
	OutVals  []string
}
