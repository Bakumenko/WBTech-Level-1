package main

import "fmt"

var (
	array2 = [5]int{2, 4, 6, 8, 10}
)

func squares2(c <-chan int) {
	for value := range c {
		sq := value * value
		fmt.Printf("Square of %v is %v\n", value, sq)
	}

}

func main() {
	fmt.Printf("Started array = %v\n", array2)
	c := make(chan int)
	go squares2(c)

	for _, val := range array2 {
		c <- val
	}
}
