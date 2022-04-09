package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) Saludar(saludo string) {
	fmt.Println(saludo, p.nombre, p.apellido)
}

func (pers persona) cumple() int {
	return pers.edad + 1
}

func main() {
	persona1 := persona{nombre: "Edgar", apellido: "Benedetto", edad: 25}
	persona1.Saludar("Hola")
	fmt.Println(persona1.cumple())
}
