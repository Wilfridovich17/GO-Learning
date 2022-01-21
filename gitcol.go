package main

import "fmt"

func main() {
	fmt.Println("Github Collaboration")
	fmt.Println("Printing digits")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Printing even digits")
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
