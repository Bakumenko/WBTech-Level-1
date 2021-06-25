package main

import "fmt"

func worker2(c chan int) {
	i := 1
	for {
		select {
		case c <- i:
			i++
		case <-c:
			return
		}
	}
}

func main() {
	c := make(chan int)
	go worker2(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	close(c)
}
