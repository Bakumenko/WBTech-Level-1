package main

import "fmt"

var (
	array11 = [5]int{2, 5, 3, 1, 4}
	array21 = [6]int{5, 7, 6, 4, 8, 3}
)

func main() {
	for _, val1 := range array11 {
		for _, val2 := range array21 {
			if val1 == val2 {
				fmt.Println(val1)
			}
		}
	}
}
