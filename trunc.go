package main
import(
	"fmt"
	"strconv"
)

func main(){
	var numero float64

	fmt.Printf("Hello, write a number: ")
	fmt.Scan(&numero)

	fmt.Printf("Your number is: %s\n", strconv.FormatFloat(numero, 'f', -1, 64))
	fmt.Printf("The truncated number is: %s\n", strconv.Itoa(int(numero)))
}