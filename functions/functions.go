package main

import "fmt"

var addFunc = func(a, b int) int {
	return a + b
}

func main() {
	// fmt.Println(addFunc(1, 2))
	// AddHigherOrderFunction(addFunc)
	closureFunc := AddClosureFunction()
	fmt.Println(closureFunc())
	fmt.Println(closureFunc())
}

func Add(a, b int) int {
	return a + b
}

func AddHigherOrderFunction(addHigherOrderFunction func(a, b int) int) {
	fmt.Println(addHigherOrderFunction(3, 2))
}

func AddClosureFunction() func() int {
	a := 4
	b := 6
	return func() int {
		a++
		b++
		return addFunc(a, b)
	}
}
