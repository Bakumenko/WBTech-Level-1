package main

import "fmt"

var (
	v1 = 1
	v2 = "oleg"
	v3 = true
	v4 = make(chan int)
	v5 = 1.0
)

func defineType(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Printf("%v is int\n", variable)
	case string:
		fmt.Printf("%v is string\n", variable)
	case bool:
		fmt.Printf("%v is bool\n", variable)
	case chan int:
		fmt.Printf("%v is chan int\n", variable)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	defineType(v1)
	defineType(v2)
	defineType(v3)
	defineType(v4)
	defineType(v5)
}
