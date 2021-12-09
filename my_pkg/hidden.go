package my_pkg

import "fmt"

var hidden_var int = 1

func Watch_hidden_var() {
	fmt.Printf("%d", hidden_var)
}
