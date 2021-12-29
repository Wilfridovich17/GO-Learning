package main

import (
	"fmt"
	"sync"
)

var i int = 0
var wg sync.WaitGroup
var mut sync.Mutex

func increase() {
	mut.Lock()
	i += 1
	mut.Unlock()
	wg.Done()
}

func main() {
	wg.Add(8)
	go increase()
	go increase()
	go increase()
	go increase()
	go increase()
	go increase()
	go increase()
	go increase()
	wg.Wait()
	fmt.Println(i)
}
