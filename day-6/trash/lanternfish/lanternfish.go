package lanternfish

import (
	"fmt"
)

type LanternFish struct {
	id          int
	value       int
	Tick        chan int
	repro       chan<- int
	Quit        chan bool
	maxLifespan int
}

func NewLanternFish(id int, value int, lifespan int, repro chan<- int) *LanternFish {
	return &LanternFish{
		id:          id,
		value:       value,
		repro:       repro,
		Tick:        make(chan int),
		Quit:        make(chan bool),
		maxLifespan: lifespan,
	}
}

// Start : starts the lantern fish listening for events with a tick channel and quit channel.  It returns
// a reproduction channel that can be used to signal the creation of a new LanternFish.
func (f *LanternFish) Start() {
	fmt.Printf("LanternFish %d is ready...\n", f.id)
	for {
		select {
		case day := <-f.Tick:
			f.processTick(day)
		case ok := <-f.Quit:
			fmt.Printf("Terminating fish %d - %t\n", f.id, ok)
			return
		}
	}
}

func (f *LanternFish) processTick(day int) {
	if day == f.maxLifespan {
		f.Quit <- true
	}
	fmt.Printf("LanternFish %d is ticking day %d!\n", f.id, day+1)
	f.value--
	if f.value < 0 {
		f.value = 6

		// reproduce
		fmt.Printf("Reproducing: %d\n", f.id)
		f.repro <- f.id
	}
	//fmt.Printf("LanternFish %d ended day %d with value %d\n", f.id, day, f.value)
}
