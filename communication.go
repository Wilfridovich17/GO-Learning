package main

import (
	"fmt"
	"sync"
)

func simple_product(i1 int, i2 int, c chan int, wg *sync.WaitGroup) {
	c <- i1 * i2
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	wg.Add(2)
	go simple_product(1, 2, c, &wg)
	go simple_product(3, 4, c, &wg)
	a := <-c
	b := <-c
	wg.Wait()
	fmt.Println(a)
	fmt.Println(b)
}
