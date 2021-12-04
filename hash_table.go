package main

import (
	"fmt"
)

func main() {
	my_hash_table := map[string]int{"John": 15, "Mike": 18}
	fmt.Println(my_hash_table["John"])

	//Create a new record (Similar to Python)
	my_hash_table["Joanna"] = 17

	//Delete a record (Similar to Python)
	delete(my_hash_table, "Mike")

	//We can use a two assignment to test the existence of a key
	value, it_exists := my_hash_table["Joanna"]
	fmt.Println(value, it_exists)
	value, it_exists = my_hash_table["Mike"]
	fmt.Println(value, it_exists)
	//We can iterate through a map (similar to Python)
	for key, val := range my_hash_table {
		fmt.Println(key, val)
	}

	fmt.Println(len(my_hash_table))

}
