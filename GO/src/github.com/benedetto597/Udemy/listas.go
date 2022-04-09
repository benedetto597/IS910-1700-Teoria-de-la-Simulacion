package main

import "fmt"

func main() {
	frutas := []string{"Manzana", "Pera", "Naranja"}

	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}

	// Añadir a la lista más elementos
	frutas = append(frutas, "Sandia", "Melon")

	// Cambiar el valor de una posición en la lista
	frutas[0] = "Platano"

	for i := 0; i < len(frutas); i++ {
		fmt.Println(i, frutas[i])
	}
}
