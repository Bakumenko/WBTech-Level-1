package main

import "fmt"

var (
	array = [5]int{2, 4, 6, 8, 10}
)

func squares(c chan int) {
	for _, value := range array {
		c <- value * value
	}

	close(c)
}

func main() {
	fmt.Printf("Started array = %v\n", array)
	c := make(chan int)
	var sum int
	go squares(c)

	for val := range c {
		sum += val
	}

	fmt.Printf("Sum = %v\n", sum)
}
