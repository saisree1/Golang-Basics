package main

import "fmt"

// mapping two structs directly, but it restricts the fields for current struct to be same as the previous

type A struct {
	Id   int
	Name string
	B    *B
}

type B struct {
	Id   int
	Name string
}

type ADomain struct {
	Id   int
	Name string
	B    *BDomain
}

type BDomain struct {
	Id   int
	Name string
}

func main() {
	a := &A{1, "Raju", &B{1, "Developer"}}
	a1 := &ADomain{1, "Raju", (*BDomain)(a.B)}
	fmt.Println(*a, *a1.B)
	fmt.Printf("%T\n", (*a1.B))
}
