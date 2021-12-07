package main

import (
	"fmt"
	"strconv"
)

func main() {
	my_array := [2]int{5, 9}
	//digits := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	my_slice := make([]int, 10)

	for i := 0; i < 10; i++ {
		my_slice[i] = i
	}

	fmt.Printf("len is % and cap is %s", len(my_slice), strconv.Itoa(cap(my_slice)))

	my_slice = append(my_slice, 11)

	my_array[0] += 10
	fmt.Printf("%s", strconv.Itoa(my_array[0]))

	for idx, value := range my_slice {
		fmt.Printf("%s -- %s\n", strconv.Itoa(idx+1), strconv.Itoa(value))
	}

	fmt.Printf("len is %d and cap is %s", len(my_slice), strconv.Itoa(cap(my_slice)))

}
