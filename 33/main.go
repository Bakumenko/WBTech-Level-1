package main

import (
	"fmt"
	"math/rand"
)

func insertInt(r chan<- int) {
	for {
		rnd := rand.Intn(100)
		r <- rnd
		if rnd == 50 { //остановка записи в канал, если рандомное число == 50
			close(r)
			return
		}
	}
}

func checkEven(r <-chan int, e chan<- int) {
	for val := range r {
		if val%2 == 0 {
			e <- val
		}
	}
	close(e)
}

func main() {
	r := make(chan int)
	e := make(chan int)

	go insertInt(r)
	go checkEven(r, e)

	for val := range e {
		fmt.Println(val)
	}
}
