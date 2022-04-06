package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Generar numeros pseudoaleatorios usando el método de producto medios
	reader := bufio.NewReader(os.Stdin)
	// Pedir la 1er semilla al usuario
	fmt.Println("Ingrese la primer semilla (X0): ")
	entry1, _ := reader.ReadString('\n')
	strSeed1 := strings.TrimRight(entry1, "\r\n") // Remove \n and \r

	seed1, err := strconv.Atoi(strSeed1)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	// Pedir la 2da semilla al usuario
	fmt.Println("Ingrese la primer semilla (X00): ")
	entry2, _ := reader.ReadString('\n')
	strSeed2 := strings.TrimRight(entry2, "\r\n") // Remove \n and \r
	if len(strSeed2) != len(strSeed1) {
		fmt.Println("Las semillas deben tener la misma longitud")
		os.Exit(3)
	}
	seed2, err := strconv.Atoi(strSeed2)
	if err != nil {
		fmt.Println(err)
	}

	// Pedir la cantidad de numeros que se desean generar
	fmt.Println("Ingrese la cantidad de numeros que desea generar: ")
	entry3, _ := reader.ReadString('\n')
	strQty := strings.TrimRight(entry3, "\r\n") // Remove \n and \r
	qty, err := strconv.Atoi(strQty)
	if err != nil {
		fmt.Println(err)
	}
	if qty < 4 {
		fmt.Println("La cantidad de numeros a generar debe ser mayor a 3")
		os.Exit(3)
	} else {

		// Generar los numeros pseudoaleatorios
		result := MeanProduct(qty, seed1, seed2)

		// Imprimir los numeros pseudoaleatorios
		fmt.Println("=================================")
		fmt.Println("Los numeros pseudoaleatorios son: ")
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
	}

}

func MeanProduct(qty int, seed1 int, seed2 int) []float64 {
	var ri []float64
	x1 := seed1
	x2 := seed2
	i := 0
	for i < qty {
		xi := x1 * x2
		size := len(strconv.Itoa(xi))
		c := Center(size, xi)
		if c == "0" || c == "" {
			fmt.Printf("No se pueden generar más números pseudoaleatorios a partir de esta semilla %d\n", xi)
			break
		}
		ix2, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		x1 = x2
		x2 = ix2
		fc, err := strconv.ParseFloat("0."+c, 8)
		ri = append(ri, fc)
		i++
	}
	return ri
}

func Center(size int, num int) string {
	lnum := len(strconv.Itoa(num))
	// Obtener el centro de la semilla
	// Limite superior
	ls := (size - lnum/2) / 2
	// Limite inferior
	li := ((size + lnum/2) / 2)

	/*
		fmt.Println("Size n: ", len(strconv.Itoa(num)))
		fmt.Println("Size: ", size)
		fmt.Println("Num: ", num)
		fmt.Printf("Limite superior: %d\n", ls)
		fmt.Printf("Limite inferior: %d\n", li)
	*/

	result := strconv.Itoa(num)[ls:li]

	contain := strings.Contains(result, "0")
	if contain {
		pos := strings.Index(result, "0")
		if pos == 0 {
			temp := result[1:]
			pos = strings.Index(temp, "0")
			if pos == 0 {
				result = strconv.Itoa(num)[ls+1 : li+1]
			}
		}
	}
	return result
	// fmt.Println("Result: ", result)
	return result
}
