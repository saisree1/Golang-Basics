package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	go Add(1, 2)
// 	time.Sleep(1)
// 	go Add(3, 4)
// 	time.Sleep(1)
// 	fmt.Println("In main routine")
// }

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go Add(1, 2)
	go Add(3, 4)
	defer AddNew(3, 9)
	go Add(3, 0)
	wg.Wait()
}

// func main() {
// 	go Add(1, 2)
// 	go Add(3, 4)
// 	// fmt.Scanln()
// }

func Add(a, b int) {
	fmt.Println(a + b)
	wg.Done()
}

func AddNew(a, b int) {
	fmt.Println(a + b)
}
