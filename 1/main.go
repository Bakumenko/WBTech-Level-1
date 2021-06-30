package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	id  int
	hum Human
}

func main() {
	hum := Human{"Oleg", 24}
	act := Action{1, hum}
	fmt.Println(act)
}
