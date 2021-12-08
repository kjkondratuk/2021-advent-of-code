package main

import (
	"fmt"
	"github.com/kjkondratuk/2021-advent-of-code/lib"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	toInt = func(n string) int {
		nr, _ := strconv.Atoi(n)
		return nr
	}
	toInts = func(n []string) []int {
		nr := make([]int, len(n))
		for x, i := range n {
			nr[x] = toInt(i)
		}
		return nr
	}
)

func main() {
	data := lib.NewDataReader("inputs/day-6-small.txt").Read()
	lines := data.([]string)
	numstr := strings.Split(lines[0], ",")
	num := toInts(numstr)
	numDaysStr := os.Getenv("NUM_DAYS")
	numDays, _ := strconv.Atoi(numDaysStr)

	fmt.Printf("Initial state: \n")
	fmt.Printf(SPrintSeries(num))

	wg := sync.WaitGroup{}
	for i := 0; i < len(num); i++ {
		wg.Add(1)
		go advanceNumberXDays(&wg, &num[i], numDays)
	}

	wg.Wait()

	fmt.Printf(SPrintSeries(num))

	fmt.Printf("Num fish: %d\n", len(num))
}

func advanceNumberXDays(wg *sync.WaitGroup, num *int, days int) {
	for i := 0; i < days; i++ {
		*num--
		if *num < 0 {
			newChild := 8
			newLifespan := days - (i + 1)
			if newLifespan > 0 {
				fmt.Printf("Create a new fish: p: %d c: %d l: %d\n", *num, newChild, newLifespan)
				wg.Add(1)
				go advanceNumberXDays(wg, &newChild, days-(i+1))
			}
			*num = 6
		}
	}
	fmt.Printf("Fish completed at: %d\n", *num)
	wg.Done()
}

func SPrintSeries(series []int) string {
	s := ""
	for i, n := range series {
		s += fmt.Sprintf("%d", n)
		if i+1 < len(series) {
			s += fmt.Sprintf(", ")
		} else {
			s += fmt.Sprintf("\n")
		}
	}
	return s
}
