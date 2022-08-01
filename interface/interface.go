package main

import "fmt"

type Person interface {
	// Details(id int, name string) string
	Details() string
}

type Employee struct {
	id   int
	name string
}

// func Details(id int, name string) string {
// 	return fmt.Sprintf("Employee Details : Id : %d, Name : %s", id, name)
// }

func (e *Employee) Details() string {
	return fmt.Sprintf("Employee Details : Id : %d, Name : %s", e.id, e.name)
}

// type CEO struct {
// 	id      int
// 	name    string
// 	empList []Employee
// }

// func (e CEO) Details() string {
// 	return fmt.Sprintf("Employee Details : Id : %d, Name : %s", e.id, e.name)
// }

func main() {
	var a Person
	emp := &Employee{
		id:   1,
		name: "Ramu",
	}
	a = emp
	fmt.Println(a.Details())
}
