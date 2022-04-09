package main

import "fmt"

func main() {
	// Declarar variables
	var entero int
	var decimal float64

	var texto string
	var booleano bool

	// Asignar valores
	entero = 10
	decimal = 10.5
	texto = "Hola mundo"
	booleano = true

	// Imprimir valores
	fmt.Printf("%v\n", entero)
	fmt.Println(entero, decimal, texto, booleano)

	// Declarar variables con tipo inferido
	var entero2 = 10

	// Abreviación de declarar Variables
	entero3 := 10

	// Imprimir valores
	fmt.Printf("%v\n", entero2)
	fmt.Println(entero2, entero3)
}
