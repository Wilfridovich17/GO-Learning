package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	fmt.Printf("Write something: ")
	
	reader := bufio.NewScanner(os.Stdin)
    reader.Scan() 
    input_string := reader.Text()
  
	input_string_lower := strings.ToLower(input_string)
	
	is_found := strings.HasPrefix(input_string_lower, "i")
	is_found = is_found && strings.Contains(input_string_lower, "a")
	is_found = is_found && strings.HasSuffix(input_string_lower, "n")

	if is_found {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}