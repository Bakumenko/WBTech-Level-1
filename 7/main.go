package main

import "fmt"

type Pair struct {
	Key   string
	Value int
}

func writeMap(c chan Pair) {
	c <- Pair{"Oleg", 24}
	c <- Pair{"Dima", 23}
	c <- Pair{"Vanya", 23}
	c <- Pair{"Kostya", 20}
	close(c)
}

func main() {
	c := make(chan Pair)
	m := make(map[string]int, 4)
	go writeMap(c)

	for val := range c {
		m[val.Key] = val.Value
	}

	fmt.Println(m)
}
