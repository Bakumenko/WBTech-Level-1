package main

import "fmt"

var (
	array2 = [5]int{2, 4, 6, 8, 10}
)

func squares2(v <-chan int, s chan<- int) {
	var sum int
	for value := range v {
		sum += value * value
	}
	s <- sum
}

func main() {
	fmt.Printf("Started array = %v\n", array2)
	v := make(chan int)
	s := make(chan int)
	go squares2(v, s)

	for _, val := range array2 {
		v <- val
	}
	close(v)

	sum := <-s
	fmt.Printf("Sum = %v\n", sum)
}
