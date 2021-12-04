package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var counter, aux int
	user_numbers := make([]int, 3)
	user_input := "a"
	for user_input != "x" {
		fmt.Printf("Introduce a number: ")
		fmt.Scan(&user_input)
		user_input = strings.ToLower(user_input)
		if user_input != "x" && counter > 2 {
			user_numbers = append(user_numbers, strconv.Atoi(user_input))
		}
		if user_input != "x" && counter <= 2 {
			user_numbers[counter] = strconv.Atoi(user_input)
		}
		
	}
	for _, user_number := range user_numbers{
		fmt.Println(user_number)
	}
}
