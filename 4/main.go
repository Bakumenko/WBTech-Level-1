package main

import (
	"fmt"
	"time"
)

var (
	count_workers = 20
)

func worker(id int, c <-chan interface{}) {
	for value := range c {
		time.Sleep(time.Millisecond)
		fmt.Printf("%v worker: %v\n", id, value)
	}
}

func main() {
	c := make(chan interface{})

	for id := 0; id < count_workers; id++ {
		go worker(id, c)
	}

	for i := 0; i < 40; i++ {
		c <- 123
		c <- "oleg"
		c <- struct {
			id   int
			name string
		}{1, "oleg"}
	}

	close(c)
}
