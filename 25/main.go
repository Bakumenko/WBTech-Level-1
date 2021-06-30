package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

func NewCounter() Counter {
	return Counter{0}
}

func (c *Counter) add() {
	c.count++
}

func (c *Counter) PrintCounter() {
	fmt.Println("Counter: ", c.count)
}

func (c *Counter) Work(i <-chan interface{}) {
	for value := range i {
		c.add()
		fmt.Printf("%v\n", value)
	}
}

func main() {
	c := make(chan interface{})
	counter := NewCounter()
	go counter.Work(c)

	c <- 123
	c <- "oleg"
	c <- struct {
		id   int
		name string
	}{1, "oleg"}

	time.Sleep(1 * time.Second)

	close(c)

	counter.PrintCounter()
}
