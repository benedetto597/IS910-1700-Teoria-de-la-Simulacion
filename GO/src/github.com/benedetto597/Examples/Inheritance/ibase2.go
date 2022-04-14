package main

import "fmt"

type iBase1 interface {
	say()
}
type iBase2 interface {
	walk()
}
type base1 struct {
	color string
}

func (b *base1) say() {
	fmt.Println("Hi from say function")
}

type base2 struct {
}

func (b *base1) walk() {
	fmt.Println("Hi from walk function")
}

type child struct {
	base1 //embedding
	base2 //embedding
	style string
}

func (b *child) clear() {
	fmt.Println("Clear from child's function")
}
func check1(b iBase1) {
	b.say()
}
func check2(b iBase2) {
	b.walk()
}
func main() {
	base1 := base1{color: "Red"}
	base2 := base2{}
	child := &child{
		base1: base1,
		base2: base2,
		style: "somestyle",
	}
	child.say()
	child.walk()
	check1(child)
	check2(child)
}
