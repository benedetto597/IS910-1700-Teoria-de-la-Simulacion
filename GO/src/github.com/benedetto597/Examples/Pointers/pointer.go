package main

import (
	"fmt"
)

func main() {
	num := 100
	var letra string
	// Un puntero sirven para acceder la dirección de memoria de una variable y poder modificarla.
	// Una variable de tipo puntero es una variable que almacena una dirección de memoria.
	// Se hace referencia a un puntero usando "&"
	fmt.Println(num)
	fmt.Println(&num)

	// Pedir datos al usuario con fmt.Scan()
	fmt.Scan(&letra)
}
