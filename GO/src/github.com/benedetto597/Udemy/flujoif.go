package main

import "fmt"

func main() {
	numero := 46
	if numero == 45 {
		fmt.Println("El numero es 45")
	} else if numero == 46 {
		fmt.Println("El numero es 46")
	} else {
		fmt.Println("El numero no es 45 ni 46")
	}

	edad := 30
	if edad >= 18 {
		fmt.Println("Es mayor de edad")
	} else if edad < 18 && edad >= 0 {
		fmt.Println("Es menor de edad")
	} else {
		fmt.Println("No es una edad valida")
	}

}
