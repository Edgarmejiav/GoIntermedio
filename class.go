package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func (e *Employee) setID(id int) {
	e.id = id
}
func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}
func (e *Employee) GetName() string {
	return e.name
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	e := Employee{}
	fmt.Println("%v", e)

	e.setID(23)
	e.setName("Edgar")
	fmt.Println(e.GetName())
	fmt.Println(e.GetId())
	fmt.Println(e)
	e2 := Employee{
		id:       1,
		name:     "Edgar",
		vacation: true,
	}

	fmt.Println(e2)

	e3 := new(Employee)
	e3.id = 3
	e3.name = "e3"
	e3.vacation = true
	fmt.Println("e3", *e3)

	e4 := NewEmployee(4, "e4", true)

	fmt.Println("e4", *e4)
}
