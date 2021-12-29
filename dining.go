package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var host = make(chan bool, 2)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	no              int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup, host_channel chan bool) {
	for i := 0; i < 3; i++ {
		aux := 0
		host_channel <- true

		if rand.Float64() < 0.5 {
			aux = 1
		}

		if aux == 0 {
			p.rightCS.Lock()
			p.leftCS.Lock()

			fmt.Println("starting to eat ", p.no)
			fmt.Println("finishing eating ", p.no)

			p.leftCS.Unlock()
			p.rightCS.Unlock()
		} else {
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Println("starting to eat ", p.no)
			fmt.Println("finishing eating ", p.no)

			p.rightCS.Unlock()
			p.leftCS.Unlock()
		}

		<-host_channel
	}
	wg.Done()
}

func init_chops() []*ChopS {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	return CSticks
}

func init_philos(CSticks []*ChopS) []*Philo {
	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	return philos

}

func main() {
	var wg sync.WaitGroup
	CSticks := init_chops()

	philos := init_philos(CSticks)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg, host)
	}
	wg.Wait()
}
