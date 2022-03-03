package main

import "fmt"

func main() {
	print_factorial(10)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func print_factorial(num int) {
	fmt.Println("============= Factorial =============")
	fmt.Print(fmt.Sprintf("El factorial de %[1]d es %[2]d", num, factorial(num)))
}
