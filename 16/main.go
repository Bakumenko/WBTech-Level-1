package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1 //из-за := это локальная n для оператора if, которая будет жить внутри {}. Чтоб вывести 2, нужно сделать просто =
		n++
	}
	fmt.Println(n)
}
