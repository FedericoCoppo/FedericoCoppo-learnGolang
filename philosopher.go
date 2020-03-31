package main

import (
	"fmt"
	"sync"
	"time"
)

// chopstic type
type ChopS struct {
	sync.Mutex
}

// philo type
type Philo struct {
	leftCS, rightCS * ChopS
	eatCnt int
	index int
	wgAlreadyDone bool
}

// global label
var wg sync.WaitGroup
var chopStIsBusy = [5]bool{false, false, false, false, false}
var chopStIsBusylMutex sync.Mutex

// method for eating
func (p * Philo) eat () {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.index)
		time.Sleep(150 * time.Millisecond) // eating time
		fmt.Printf("finishing eating %v\n", p.index)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		p.eatCnt = p.eatCnt + 1
		//fmt.Printf("Ph %d end to eat for the %d time \n", p.index + 1, p.eatCnt)
}

// host method
func (p * Philo) askEatToHost() {
	for {
		// manage the end of the task
		 if p.eatCnt >= 3  {
			 if p.wgAlreadyDone == false {
				 p.wgAlreadyDone = true
				 wg.Done()
			 }
		}else
		{
			chopStIsBusylMutex.Lock()

			// ALGO: eat allowed only if both chopstick are available
			if chopStIsBusy[p.index] == false && chopStIsBusy[(p.index+1)%5] == false {

				// no need to check if more that 2 philo are eating (since if a couple is available at least only 1 pholo is eating)
				chopStIsBusy[p.index] = true
				chopStIsBusy[(p.index+1)%5] = true
				chopStIsBusylMutex.Unlock()

				p.eat()

				chopStIsBusylMutex.Lock()
				chopStIsBusy[p.index] = false
				chopStIsBusy[(p.index+1)%5] = false
				chopStIsBusylMutex.Unlock()
				} else {
					chopStIsBusylMutex.Unlock()
				}
		}
	}
}

// main routine
func main() {

	// init
	Cs := make([]* ChopS, 5)

	// mutex creation
	for i:= 0; i < 5; i++ {
		Cs [i] = new(ChopS)
	}

	// philopher creation
	philos := make([]* Philo,5)

	for i:= 0; i < 5; i++ {
		philos[i] = &Philo {Cs[i],Cs[(i + 1)%5], 0 , i , false }
	}

	// sync with main
	wg.Add(5)

	for i:= 0; i < 5; i++ {
		go philos[i].askEatToHost()
	}

	// to end main, wait all 5 eat
	wg.Wait()
	fmt.Println("All philosopher eat 3 times, main end..")
}
