package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//ANIMAL INTERFACE
type Animal interface {
	Eat()
	Move()
	Speak()
}

func print_info(a Animal, info string) {
	switch info {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

//Function to create animals
func Create_Animal(animal_type string) Animal {
	var user_animal Animal

	switch animal_type {
	case "cow":
		user_animal = Cow{food: "grass", locomotion: "walk", noise: "moo"}
	case "snake":
		user_animal = Snake{food: "mice", locomotion: "slither", noise: "hsss"}
	case "bird":
		user_animal = Bird{food: "worms", locomotion: "fly", noise: "peep"}
	}
	return user_animal
}

func validate_command(command string) (string, bool) {
	aux := strings.ToLower(command)
	if (aux == "query") || (aux == "newanimal") {
		return aux, true
	} else {
		return aux, false
	}
}

func validate_type(command string) (string, bool) {
	aux := strings.ToLower(command)
	if (aux == "cow") || (aux == "bird") || (aux == "snake") {
		return aux, true
	} else {
		return aux, false
	}
}

func validate_info(command string) (string, bool) {
	aux := strings.ToLower(command)
	if (aux == "eat") || (aux == "move") || (aux == "speak") {
		return aux, true
	} else {
		return aux, false
	}
}

//TYPE COW AND ITS METHODS
type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() {
	fmt.Printf("%s\n", a.food)
}

func (a Cow) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

func (a Cow) Speak() {
	fmt.Printf("%s\n", a.noise)
}

//TYPE BIRD AND ITS METHODS
type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (a Bird) Eat() {
	fmt.Printf("%s\n", a.food)
}

func (a Bird) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

func (a Bird) Speak() {
	fmt.Printf("%s\n", a.noise)
}

//TYPE SNAKE AND ITS METHODS
type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (a Snake) Eat() {
	fmt.Printf("%s\n", a.food)
}

func (a Snake) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

func (a Snake) Speak() {
	fmt.Printf("%s\n", a.noise)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	user_animals := make(map[string]Animal)

	user_animals["my_animal"] = Create_Animal("cow")
	user_animals["my_animal"].Speak()
	for {
		fmt.Printf("Example request: <newanimal> <name> <type> or <query> <name> <info>.\nPress x to exit.\n")
		fmt.Printf("> ")
		reader.Scan()

		if strings.ToLower(reader.Text()) == "x" {
			break
		}

		user_input := strings.Split(reader.Text(), " ")
		if len(user_input) == 3 {
			cmd, val := validate_command(user_input[0])
			if (cmd == "newanimal") && (val == true) {
				atype, val := validate_type(user_input[2])
				if val == true {
					user_animals[user_input[1]] = Create_Animal(atype)
					fmt.Println("Created it!")
				} else {
					fmt.Println("Invalid animal type")
				}
			}
			if (cmd == "query") && (val == true) {
				info, val := validate_info(user_input[2])
				if val == true {
					print_info(user_animals[user_input[1]], info)
				} else {
					fmt.Println("Data no available")
				}
			}
		} else {
			fmt.Println("Invalid command")
		}
	}
}
