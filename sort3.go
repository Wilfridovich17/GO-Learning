package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func Swap(list_numbers []int, idx int) {
	var bubble int
	bubble = list_numbers[idx]
	list_numbers[idx] = list_numbers[idx+1]
	list_numbers[idx+1] = bubble
}

func Sort(list_numbers []int, wg *sync.WaitGroup, nb int) {
	var is_swap bool = true
	for i := 0; i < len(list_numbers)-1; i++ {
		is_swap = false
		for j := 0; j < len(list_numbers)-1; j++ {
			if list_numbers[j] > list_numbers[j+1] {
				Swap(list_numbers, j)
				is_swap = true
			}
		}
		if is_swap == false {
			break
		}
	}
	if nb > 0 {
		fmt.Println("Block no. ", nb, list_numbers)
		wg.Done()
	}
}

func main() {
	var user_input, split_size, counter, ul, ite int
	var wg sync.WaitGroup
	reader := bufio.NewScanner(os.Stdin)
	user_numbers := make([]int, 0)

	for {
		fmt.Printf("Write a number or press x to end: ")
		reader.Scan()
		if strings.ToLower(reader.Text()) == "x" {
			break
		} else {
			user_input, _ = strconv.Atoi(reader.Text())
			user_numbers = append(user_numbers, user_input)
		}
	}

	for j := 4; j >= 1; j = j / 2 {
		fmt.Println("Iteration ", ite+1)
		if len(user_numbers)%j > 0 {
			split_size = len(user_numbers)/j + 1
		} else {
			split_size = len(user_numbers) / j
		}
		/*We create 4 go routines, depending of the numbers introduced by the user
		for example, if the user gives 6 numbers, we create 3 blocks of 2 numbers and 1 empty block
		*/
		for i := 0; i < len(user_numbers) && i < j; i++ {
			wg.Add(1)
			ul = int(math.Min(float64(len(user_numbers)), float64(split_size*(i+1))))
			go Sort(user_numbers[(split_size*i):ul], &wg, counter+1)
			counter++
		}
		wg.Wait()
		ite++
	}
	fmt.Println("Final block: ", user_numbers)

}
