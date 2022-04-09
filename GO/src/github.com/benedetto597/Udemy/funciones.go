package main

import "fmt"

func main() {
	Saludar("Edgar")
	Sumar(10, 20)
	suma := Sumar(5, 10)
	fmt.Println(suma)
	NumerosPares(15)
}

func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func Sumar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func NumerosPares(numero int) {
	for i := 0; i <= numero; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
