package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sort(slice []int, new_number int) int {
	for idx, value := range slice {
		if value > new_number {
			return (idx)
		}
	}
	return (-1)
}

func main() {
	var aux, counter, cut int
	var err error
	var user_input string = "a"

	user_numbers := make([]int, 3)

	for user_input != "x" {
		fmt.Printf("Introduce a number: ")
		fmt.Scan(&user_input)
		aux, err = strconv.Atoi(user_input)
		if err == nil {
			switch {
			case counter == 0:
				user_numbers[counter] = aux
			case counter > 0:
				cut = sort(user_numbers, aux)
				user_numbers = append(user_numbers[0:cut+1], user_numbers[cut:counter]...)
				user_numbers[cut] = aux
			}
			counter++
		} else {
			user_input = strings.ToLower(user_input)
		}
	}
	for _, user_number := range user_numbers {
		fmt.Println(user_number)
	}
}
