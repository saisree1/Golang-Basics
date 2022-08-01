package main

import "fmt"

func main() {
	var a *string
	fmt.Println(a)
	aTemp := "Ramu"
	a = &aTemp
	fmt.Println(a, *a)

	// best way
	var b string = "Ramu"
	fmt.Println(&b, b)

	emp := &Employee{
		id:   1,
		name: "Ramu",
	}
	fmt.Println(emp, *emp, emp.id, emp.name)

	empWithPointers := &EmployeeWithPointers{
		id:   &emp.id,
		name: &emp.name,
	}
	fmt.Println(empWithPointers, *empWithPointers, *empWithPointers.id, *empWithPointers.name)
}

type Employee struct {
	id   int
	name string
}

type EmployeeWithPointers struct {
	id   *int
	name *string
}
