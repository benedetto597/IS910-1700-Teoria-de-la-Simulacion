package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	// Crear objetos a partir de la estructura
	p1 := persona{nombre: "Edgar", apellido: "Benedetto", edad: 25}
	p2 := persona{nombre: "Juan", apellido: "Perez", edad: 30}

	// Imprimir valores
	fmt.Println("Persona 1: ", p1)
	fmt.Println("Nombre Persona 1: ", p1.nombre)
	fmt.Println("Persona 2: ", p2)
	fmt.Println("Apellido Persona 2: ", p2.apellido)

	// Declarar variables con tipo inferido
	var p3 persona
	p3.nombre = "Pedro"
	p3.apellido = "Perez"
	p3.edad = 25

	// Imprimir valores
	fmt.Println("Persona 3: ", p3)

	// Cambiar los valores de un objeto a partir de la estructura
	p3.nombre = "Juan"
	p3.apellido = "Godoy"
	p3.edad = 30

	// Imprimir valores
	fmt.Println("Persona 3: ", p3)
}
