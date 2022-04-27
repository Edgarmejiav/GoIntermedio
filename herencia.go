package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee1 struct {
	id int
}

// composición
type FullTimeEmployee struct {
	Person
	Employee1
	endDate int
}
type TemporaryEmployee struct {
	Person
	Employee1
	taxRate int
}
type PrinteInfo interface {
	getMenssage() string
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

func (tfEmployee FullTimeEmployee) getMessage() string {
	return "Full Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {

	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
