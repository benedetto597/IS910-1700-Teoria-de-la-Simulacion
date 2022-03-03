package main

import (
	"fmt"

	"github.com/benedetto597/Exercises/Factorial"

	"github.com/benedetto597/Exercises/String_To_Date"
)

func main() {
	fmt.Println("=============== Convert String to Date ===============")
	String_To_Date.Convert()
	fmt.Println("-------------------------------------")
	fmt.Println("\n=============== Factorial de un número ===============")
	Factorial.Calculate()
}
