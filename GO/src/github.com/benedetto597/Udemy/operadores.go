package main

import "fmt"

func main() {
	// Operadores matematicos
	suma := 10 + 5
	resta := 30 - 5
	multiplicacion := 10 * 5
	division := 10 / 5
	modulo := 20 % 5

	// Imprime el resultado de las operaciones
	fmt.Println(suma, resta, multiplicacion, division, modulo)

	suma++
	suma += 5
	resta--
	resta -= 5
	multiplicacion *= 2
	division /= 2
	modulo %= 2

	// Imprime el resultado de las operaciones
	fmt.Println(suma, resta, multiplicacion, division, modulo)

	// Operadores logicos
	mayorque := 10 > 5
	menorque := 10 < 5
	mayorigualque := 10 >= 5
	menorigualque := 10 <= 5
	igualque := 10 == 5
	diferenteque := 10 != 5

	operadory := 10 > 5 && 10 < 20
	operadoro := 10 > 5 || 10 < 5

	// Imprime el resultado de las operaciones
	fmt.Println(mayorque, menorque, mayorigualque, menorigualque, igualque, diferenteque, operadory, operadoro)
}
