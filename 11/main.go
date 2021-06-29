package main

import "fmt"

var (
	slice11 = []int{2, 5, 3, 1, 4}
	slice21 = []int{5, 7, 6, 4, 8, 3}
)

func crossingSlices1(s1 []int, s2 []int) []int {
	var cross []int
	for _, val1 := range s1 {
		for _, val2 := range s2 {
			if val1 == val2 {
				cross = append(cross, val1)
			}
		}
	}
	return cross
}

func main() {
	cross := crossingSlices1(slice11, slice21)
	fmt.Println("Slice1: ", slice11)
	fmt.Println("Slice2: ", slice21)
	fmt.Println("Crossing: ", cross)
}
