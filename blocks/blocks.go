package main

import "fmt"

func main() {
	var a []int
	a = append(a, 10)

	for ind, val := range a {
		fmt.Println(ind, val)
	}

	var b2 []int
	b2 = append(b2, 20)

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	// for {
	// 	fmt.Println("while loop")
	// }

	var b int
	b = 1

	if b == 0 {
		fmt.Println("Initialised")
	} else if b == 1 {
		fmt.Println("updated")
	} else {
		fmt.Println("default value")
	}

	switch b {
	case 0:
		fmt.Println("Initialised")
		break
	case 1:
		fmt.Println("updated")
	default:
		fmt.Println("default value")
	}
}
