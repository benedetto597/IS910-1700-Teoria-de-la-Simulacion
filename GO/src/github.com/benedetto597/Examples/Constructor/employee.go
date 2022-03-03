package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

func (e *Employee) Init(name string, age int) {
	e.Name = name
	e.Age = age
}

func info(name string, age int) *Employee {
	e := new(Employee)
	e.Name = name
	e.Age = age
	return e
}

func main() {
	emp := new(Employee)
	emp.Init("John Doe", 25)
	fmt.Printf("%s: %d\n", emp.Name, emp.Age)

	empInfo := info("John Doe", 25)
	fmt.Printf("%v", empInfo)
}
