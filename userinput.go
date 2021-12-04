package main
import (
	"fmt"
	"strconv"
)

func main(){
	var age int
	fmt.Printf("Escribe tu edad ")
	//num, err := fmt.Scan(&age)
	fmt.Scan(&age)

	fmt.Printf("Tu edad es: %s\n", strconv.Itoa(age))
	//fmt.Printf("Se escanearon %s elementos y el código de error es %s\n", strconv.Itoa(num), err)
}