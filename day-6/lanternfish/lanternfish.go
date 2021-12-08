package lanternfish

import (
	"fmt"
)

type lanternFish struct {
	id    int
	value int
	tick  chan int
	repro chan bool
	//quit  chan bool
}

type LanternFish interface {
	Send(int)
	Start()
	//Stop()
}

func NewLanternFish(id int, value int, repro chan bool) LanternFish {
	return &lanternFish{
		id:    id,
		value: value,
		//quit:  make(chan bool),
		repro: repro,
		tick:  make(chan int),
	}
}

// Start : starts the lantern fish listening for events with a tick channel and quit channel.  It returns
// a reproduction channel that can be used to signal the creation of a new LanternFish.
func (f *lanternFish) Start() {
	fmt.Printf("LanternFish %d is ready...\n", f.id)
	for {
		select {
		case day := <-f.tick:
			f.processTick(day)
		//case <-f.quit:
		//	fmt.Printf("Stopping fish: %d end value is %d\n", f.id, f.value)
		//	return
		}
	}
}

//func (f *lanternFish) Stop() {
//	f.quit <- true
//}

func (f *lanternFish) Send(d int) {
	f.tick <- d
}

func (f *lanternFish) processTick(day int) {
	fmt.Printf("LanternFish %d is ticking day %d!\n", f.id, day+1)
	f.value--
	if f.value < 0 {
		f.value = 6

		// reproduce
		//wg.Add(1)
		f.repro <- true
		fmt.Printf("Done reproducing\n")
	}
	fmt.Printf("LanternFish %d ended day %d with value %d\n", f.id, day, f.value)
}
