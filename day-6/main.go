package main

import (
	"fmt"
	"github.com/kjkondratuk/2021-advent-of-code/day-6/lanternfish"
	"github.com/kjkondratuk/2021-advent-of-code/lib"
	"os"
	"strconv"
	"strings"
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

	fmt.Printf("Initial state: ")
	fmt.Printf(SPrintSeries(num))

	fish := make([]lanternfish.LanternFish, len(num))
	tick := make(chan int)
	repro := make(chan bool)
	//stopAll := make(chan bool)

	for i, n := range num {
		f := lanternfish.NewLanternFish(i, n, repro)
		fish[i] = f
		//fwg.Add(1)
		fmt.Printf("Starting fish: %d\n", i)
		go f.Start()
	}

	subsequentSpawns := 0
	go func() {
		for {
			select {
			case day := <-tick:
				for _, f := range fish {
					f.Send(day)
				}
			case _ = <-repro:
				fmt.Printf("Spawning new fish: %d\n", subsequentSpawns)
				subsequentSpawns++
				f := lanternfish.NewLanternFish(0, 8, repro)
				fish = append(fish, f)
				//fwg.Add(1)
				go f.Start()
			//case <-stopAll:
			//	for _, f := range fish {
			//		f.Stop()
			//	}
			}
		}
	}()

	fmt.Printf("Initial fish started...\n")

	for i := 0; i < numDays; i++ {
		tick <- i
	}

	// send stop signal
	//stopAll <- true

	fmt.Printf("Done simulating...")

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
