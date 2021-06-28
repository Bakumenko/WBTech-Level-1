package main

import "fmt"

func outputSlice2(c chan int, s <-chan []int) {
	slice := <-s
	for _, value := range slice {
		c <- value
	}

	close(c)
}

func main() {
	slice2 := []int{1, 2, 4, 6, 8, 10}
	fmt.Printf("Started slice = %v\n", slice2)
	c := make(chan int)
	s := make(chan []int)

	go outputSlice2(c, s)

	s <- slice2

	for val := range c {
		fmt.Printf("Value = %v\n", val)
	}

}
