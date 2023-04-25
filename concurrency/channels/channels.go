package main

import (
	"time"
)

// func CalculateValue(values chan int) {
// 	value := rand.Intn(10)
// 	fmt.Println("Calculated Random Value")
// 	values <- value
// }

// func main() {
// 	rand.Seed(time.Now().Unix())

// 	values := make(chan int, 2)
// 	defer close(values)

// 	go CalculateValue(values)
// 	// fmt.Scanln()
// 	value := <-values
// 	fmt.Println(value)
// }

// func access(ch chan int) {
// 	time.Sleep(time.Second)
// 	fmt.Println("start accessing channel\n")

// 	for i := range ch {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	// only modify this line to define the capacity
// 	ch := make(chan int, 3)
// 	defer close(ch)

// 	go access(ch)

// 	for i := 0; i < 9; i++ {
// 		ch <- i
// 		fmt.Println("Filled: ", i)
// 	}

// 	time.Sleep(3 * time.Second)
// }

func main() {
	ch := make(chan int, 1)
	defer close(ch)

	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)

	ch <- 1
	ch <- 2
}
