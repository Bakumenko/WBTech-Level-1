package main

import "fmt"

var (
	first2  = 3
	second2 = 5
)

func main() {
	fmt.Printf("First number = %v\nSecond number = %v\n\n", first2, second2)

	first2 = first2 + second2
	second2 = first2 - second2
	first2 = first2 - second2

	fmt.Printf("After switch:\nFirst number = %v\nSecond number = %v\n", first2, second2)
}
