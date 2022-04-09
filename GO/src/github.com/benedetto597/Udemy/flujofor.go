package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("==========================")

	nombre := "Edgar"

	for i := 0; i < len(nombre); i++ {
		// Formatear a texto usando string
		fmt.Println(i, string(nombre[i]))
	}
	fmt.Println("==========================")

	// En go no existe el while por lo que se usa for
	numero := 5
	for numero != 1 {
		fmt.Println(numero)
		numero--
	}
	fmt.Println("==========================")
}
