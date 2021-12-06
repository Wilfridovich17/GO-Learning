package main

import (
	"bufio"
	"fmt"
	"os"
)

func position(element []byte) int64 {
	for idx, val := range element {
		if (val == 10) || (val == 32) {
			return (int64(idx))
		}
	}
	return (20)
}

func main() {
	type Persona struct {
		name     string
		lastname string
	}

	var aux Persona
	var counter int
	var cut, c_pos, n_pos, step int64
	var err_2 error

	people := make([]Persona, 0)
	read_line := make([]byte, 20)

	reader := bufio.NewScanner(os.Stdin)

	fmt.Printf("Write the name of your file: my_file.txt: ")
	reader.Scan()

	f, err_1 := os.Open(reader.Text())
	if err_1 == nil {
		for err_2 == nil {
			counter++
			_, err_2 = f.Read(read_line)
			if err_2 == nil {
				n_pos, _ = f.Seek(0, 1)
				step = n_pos - c_pos

				if step < 20 {
					for i := n_pos - c_pos; i < 20; i++ {
						read_line[i] = 0
					}
				}

				cut = position(read_line)
				if counter%2 == 0 {
					aux.lastname = string(read_line[0:cut])
					people = append(people, aux)
				} else {
					aux.name = string(read_line[0:cut])
				}

				c_pos, _ = f.Seek(+cut+1-step, 1)
			}
		}
		f.Close()
	}
	fmt.Printf("These are the names in the file:\n\n")
	for idx, value := range people {
		fmt.Printf("%d. Name: %s -- Lastname: %s\n", idx+1, value.name, value.lastname)
	}
}
