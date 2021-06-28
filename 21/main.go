package main

import "fmt"

var (
	slice = []int{1, 2, 4, 6, 8, 10}
)

func outputSlice(c chan int) {
	for _, value := range slice {
		c <- value
	}

	close(c)
}

func main() {
	fmt.Printf("Started slice = %v\n", slice)
	c := make(chan int)

	go outputSlice(c)

	for val := range c {
		fmt.Printf("Value = %v\n", val)
	}

}
