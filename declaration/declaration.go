package main

import (
	"fmt"
)

var a int

func main() {
	// var a int
	// b := 0
	// fmt.Println(a, b)

	// fmt.Println(fmt.Sprintf("Hello with id : %v", 2))

	// type IntType int
	// var a1 IntType
	// fmt.Println(a1)

	// a2 := new(IntType)
	// fmt.Println(a2, *a2)

	// var emp Employee
	// emp1 := Employee{}
	// emp2 := new(Employee)
	// fmt.Println(emp, emp1, emp2)

	// var empWithPointers *Employee
	// empWithPointers1 := &Employee{}
	// empWithPointers2 := new(Employee)
	// fmt.Println(empWithPointers, empWithPointers1, empWithPointers2)

	// var newMap map[string]string
	// newMap1 := map[string]string{}
	// newMap2 := make(map[string]string)
	// fmt.Println(newMap, newMap1, newMap2)

	// var newMapWithPointers *map[string]string
	// newMapWithPointers1 := &map[string]string{}
	// fmt.Println(newMapWithPointers, newMapWithPointers1)

	// var arr [10]int
	// fmt.Println(arr, len(arr), cap(arr))

	var arrAsSlice []int
	fmt.Println(arrAsSlice, len(arrAsSlice), cap(arrAsSlice))
}

type Employee struct {
	id   int
	name string
}
