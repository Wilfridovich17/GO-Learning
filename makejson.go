package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

func main(){
	my_map := make(map[string]string)

	reader:= bufio.NewScanner(os.Stdin)


	fmt.Printf("Write your name: ")
	reader.Scan()
	my_map["name"] = reader.Text()
	
	fmt.Printf("Write your address: ")
	reader.Scan()
	my_map["address"] = reader.Text()

	byte_rep, _ := json.Marshal(my_map)
	
	fmt.Printf("This is the JSON []byte representation:\n")
	fmt.Println(byte_rep)

}