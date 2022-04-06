package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Generar numeros pseudoaleatorios usando el método de cuadrados medios
	reader := bufio.NewReader(os.Stdin)
	// Pedir la semilla al usuario
	fmt.Println("Ingrese la semilla: ")
	entry1, _ := reader.ReadString('\n')
	strSeed := strings.TrimRight(entry1, "\r\n") // Remove \n and \r

	seed, err := strconv.Atoi(strSeed)
	if err != nil {
		fmt.Println(err)
	}

	// Pedir la cantidad de numeros que se desean generar
	fmt.Println("Ingrese la cantidad de numeros que desea generar: ")
	entry2, _ := reader.ReadString('\n')
	strQty := strings.TrimRight(entry2, "\r\n") // Remove \n and \r
	qty, err := strconv.Atoi(strQty)
	if err != nil {
		fmt.Println(err)
	}
	if qty < 4 {
		fmt.Println("La cantidad de numeros a generar debe ser mayor a 3")
	} else {

		// Generar los numeros pseudoaleatorios
		result := MeanSquare(qty, seed)

		// Imprimir los numeros pseudoaleatorios
		fmt.Println("=================================")
		fmt.Println("Los numeros pseudoaleatorios son: ")
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
	}

}

func MeanSquare(qty int, seed int) []float64 {
	size := len(strconv.Itoa(seed))
	xi := seed
	var ri []float64
	i := 0
	for i < qty {
		xi = xi * xi
		c := Center(size, xi)
		// fmt.Println("Xi: ", xi)
		// fmt.Println("C: ", c)
		ixi, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		xi = ixi
		fc, err := strconv.ParseFloat("0."+c, 8)
		ri = append(ri, fc)
		i++
	}
	return ri
}

func Center(size int, num int) string {
	exc := len(strconv.Itoa(num)) / 2
	if exc%2 != 0 {
		exc--
	}
	// Obtener el centro de la semilla
	// Limite superior
	ls := exc / 2
	// Limite inferior
	li := size + ls

	result := strconv.Itoa(num)[ls:li]

	// Si el centro inicia con 2 ceros moverse una posición
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
}
