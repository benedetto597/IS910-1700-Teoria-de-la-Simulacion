package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Seleccionar entre el metodo congruencial multiplicativo o mixto
	menu :=
		`
	Bienvenido, seleccione el método que desea usar para generar los numeros pseudoaleatorios:
	[ 1 ] Congruencial Multiplicativo
	[ 2 ] Congruencial Mixto
	[ 3 ] Salir
	`
	fmt.Print(menu)
	reader := bufio.NewReader(os.Stdin)

	entry, _ := reader.ReadString('\n')
	choose := strings.TrimRight(entry, "\r\n") // Remove \n and \r

	// Pedir la semilla al usuario
	fmt.Println("Ingrese la semilla: ")
	entrySeed, _ := reader.ReadString('\n')
	strSeed := strings.TrimRight(entrySeed, "\r\n") // Remove \n and \r

	seed, err := strconv.Atoi(strSeed)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// Pedir el multiplicador al usuario
	fmt.Println("Ingrese el multiplicador: ")
	entryMultiply, _ := reader.ReadString('\n')
	strMultiply := strings.TrimRight(entryMultiply, "\r\n") // Remove \n and \r

	multiply, err := strconv.Atoi(strMultiply)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// Pedir el modulo al usuario
	fmt.Println("Ingrese el modulo: ")
	entryModule, _ := reader.ReadString('\n')
	strModule := strings.TrimRight(entryModule, "\r\n") // Remove \n and \r

	module, err := strconv.Atoi(strModule)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	var result []float64

	switch choose {
	case "1":
		result = CongruentialMultiply(seed, multiply, module)
	case "2":
		// Pedir el modulo al usuario
		fmt.Println("Ingrese la constante: ")
		entryConstant, _ := reader.ReadString('\n')
		strConstant := strings.TrimRight(entryConstant, "\r\n") // Remove \n and \r

		constant, err := strconv.Atoi(strConstant)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		result = CongruentialMixed(seed, multiply, constant, module)

	case "3":
		os.Exit(0)
	default:
		log.Fatalf("Selecciono una opción que no es valida")
		os.Exit(3)
	}

	// Imprimir los numeros pseudoaleatorios
	fmt.Println("=================================")
	fmt.Println("Los numeros pseudoaleatorios son: ")
	for i := 0; i < len(result); i++ {
		fmt.Printf("%.4f\n", result[i])
	}

}

func CongruentialMultiply(seed int, multiply int, module int) []float64 {
	centinel := false
	var ri []float64
	for centinel != true {
		seed = (seed * multiply) % module
		ui := float64(seed) / (float64(module) - 1)
		alreadyExist := ExistInArray(ri, ui)
		if alreadyExist {
			centinel = true
		} else {
			ri = append(ri, ui)
		}
	}
	return ri
}

func CongruentialMixed(seed int, multiply int, constant int, module int) []float64 {
	centinel := false
	var ri []float64
	for centinel != true {
		seed = ((seed * multiply) + constant) % module
		ui := float64(seed) / float64(module)
		alreadyExist := ExistInArray(ri, ui)
		if alreadyExist {
			centinel = true
		} else {
			ri = append(ri, ui)
		}
	}
	return ri
}

func ExistInArray(arr []float64, value float64) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
