package main

import (
	"fmt"
	"time"
)

var countSeconds = 2

func send(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func print(c chan<- int) {
	var i int
	for {
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)
	go print(c)
	go send(c)

	time.Sleep(time.Duration(countSeconds) * time.Second)
	close(c)
}
