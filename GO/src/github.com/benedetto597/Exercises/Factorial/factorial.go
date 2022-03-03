package Factorial

import "fmt"

func Calculate() {
	PrintFactorial(10)
}

func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * Factorial(num-1)
}

func PrintFactorial(num int) {
	fmt.Println("============= Factorial =============")
	fmt.Print(fmt.Sprintf("El factorial de %[1]d es %[2]d", num, Factorial(num)))
}
