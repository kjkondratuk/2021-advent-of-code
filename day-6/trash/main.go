package main

import (
	"fmt"
	"github.com/kjkondratuk/2021-advent-of-code/day-6/trash/lanternfish"
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

	fmt.Printf("Initial state: ")
	fmt.Printf(SPrintSeries(num))

	wg := sync.WaitGroup{}
	fish := make([]lanternfish.LanternFish, len(num))
	tick := make(chan int)
	repro := make(chan int)

	for i, n := range num {
		f := lanternfish.NewLanternFish(i, n, numDays, repro)
		fish[i] = *f
		fmt.Printf("Starting fish: %d\n", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			f.Start()
		}()
	}

	quit := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		orchestrate(&wg, tick, repro, numDays, fish, quit)
	}()

	fmt.Printf("Initial fish started...\n")

	for i := 0; i < numDays; i++ {
		tick <- i
	}

	wg.Wait()
	// halt orchestrator
	quit <- true

	// send stop signal
	//stopAll <- true

	fmt.Printf("Done simulating...\n")

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

func orchestrate(wg *sync.WaitGroup, tick <-chan int, repro chan int, newLifespan int, fish []lanternfish.LanternFish,
	quit <-chan bool) {
	for {
		select {
		case day := <-tick:
			for _, f := range fish {
				f.Tick <- day
			}
		case p := <-repro:
			fmt.Printf("Spawning new fish from parent: %d\n", p)
			f := lanternfish.NewLanternFish(0, 8, newLifespan, repro)
			fish = append(fish, *f)
			wg.Add(1)
			go func() {
				f.Start()
			}()
		case b := <-quit:
			fmt.Printf("Quitting orchestrator... %t\n", b)
			return
		}
	}
}
