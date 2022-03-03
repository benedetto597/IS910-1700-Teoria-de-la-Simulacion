package main

import "fmt"

func main() {
	var name string // Definir variable

	name = "Benedetto"

	fmt.Println("Hello, ", name)

	name = "Edgar"

	fmt.Println("Hello, ", name)

	var a, b int = 1, 2
	var c, d int

	c = 3
	d = 4

	fmt.Println(a, b, c, d)

	var (
		pi      float64
		boolean bool
		string  = "Text 01"
		age     = 25
	)

	pi = 3.1416
	boolean = true

	fmt.Println(pi, boolean, string, age)

	v1 := 24 // Shortcout declarar variable y asignar valor

}
