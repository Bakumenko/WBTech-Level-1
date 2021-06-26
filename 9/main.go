package main

import "fmt"

var (
	array3 = [5]int{2, 4, 6, 8, 10}
)

func double(v <-chan int, d chan<- int) {
	for value := range v {
		d <- 2 * value
	}
}

func main() {
	fmt.Printf("Started array = %v\n", array3)
	v := make(chan int)
	d := make(chan int)

	go double(v, d)

	for _, val := range array3 {
		v <- val
		fmt.Printf("Value = %v; 2 * Value = %v\n", val, <-d)
	}
	close(v)
}
