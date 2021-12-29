package main

import (
	"fmt"
	"sync"
)

func nr(wg *sync.WaitGroup) {
	fmt.Printf("New routine")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go nr(&wg)
	wg.Wait()
	fmt.Printf("Main routine")
}
