package main

import "fmt"

var (
	array3 = [5]int{2, 4, 6, 8, 10}
)

func squares3(v <-chan int, s chan<- int) {
	for value := range v {
		s <- value * value
	}
}

func main() {
	fmt.Printf("Started array = %v\n", array3)
	v := make(chan int)
	s := make(chan int)
	var sum int
	go squares3(v, s)

	for _, val := range array3 {
		v <- val
		sum += <-s
	}
	close(v)

	fmt.Printf("Sum = %v\n", sum)
}
