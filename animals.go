package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Printf("%s\n", a.food)
}

func (a Animal) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Printf("%s\n", a.noise)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	data_animals := make(map[string]Animal)
	data_animals["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	data_animals["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	data_animals["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		fmt.Printf("Example request: cow food. Press x to exit.\n")
		fmt.Printf("> ")
		reader.Scan()

		if strings.ToLower(reader.Text()) == "x" {
			break
		}

		user_input := strings.Split(reader.Text(), " ")
		switch strings.ToLower(user_input[1]) {
		case "eat":
			fmt.Printf("%s eats ", user_input[0])
			data_animals[strings.ToLower(user_input[0])].Eat()
		case "move":
			fmt.Printf("%s moves ", user_input[0])
			data_animals[strings.ToLower(user_input[0])].Move()
		case "speak":
			fmt.Printf("%s makes sound ", user_input[0])
			data_animals[strings.ToLower(user_input[0])].Speak()
		}
	}
}
