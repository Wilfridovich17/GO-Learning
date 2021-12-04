package main

import "fmt"

func main() {
	//user_numbers := make([]int, 3)
	user_input := "A"
	for user_input != "X" {
		fmt.Printf("Introduce a number: ")
		fmt.Scan(&user_input)
		fmt.Printf(user_input)
	}
}
