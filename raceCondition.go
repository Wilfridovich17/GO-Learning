
/* Explanation
------------------------------------------------------------------------------------------------------
Race Condition
A data race happens when two goroutines access the same variable concurrently, and at least one of the
access is a write instruction.
*/

package main

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup
var Counter = 0

func main() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go RoutineOne(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func RoutineOne(id int) {
	for count := 0; count < 2; count++ {
		value := Counter
		//time.Sleep(1 * time.Nanosecond) // Uncomment this line for fix race condition issue
		value++
		Counter = value
	}
	Wait.Done()
}
