package main

import "fmt"

var (
	first  = 3
	second = 5
)

func main() {
	fmt.Printf("First number = %v\nSecond number = %v\n\n", first, second)

	first = first + second
	second = first - second
	first = first - second

	fmt.Printf("After switch:\nFirst number = %v\nSecond number = %v\n", first, second)
}
