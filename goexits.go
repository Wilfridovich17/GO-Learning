package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Printf("New routine")
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Main routine")
}
