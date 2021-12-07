package main

import "fmt"

func Swap(list_numbers []int, idx int) {
	var bubble int
	bubble = list_numbers[idx]
	list_numbers[idx] = list_numbers[idx+1]
	list_numbers[idx+1] = bubble
}

func Sort(list_numbers []int) {
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
}

func main() {
	var user_input int
	user_numbers := make([]int, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("Write a number: ")
		fmt.Scan(&user_input)
		user_numbers = append(user_numbers, user_input)
	}
	Sort(user_numbers)
	fmt.Printf("Your numbers were sorted:\n")
	for _, value := range user_numbers {
		fmt.Printf("%d\t", value)
	}
}
