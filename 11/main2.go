package main

import "fmt"

var (
	slice12 = []int{2, 5, 3, 1, 4}
	slice22 = []int{5, 7, 6, 4, 8, 3}
)

func crossingSlices2(s1 []int, s2 []int) []int {
	m := make(map[int]bool)
	for _, val := range s1 {
		if _, ok := m[val]; ok {
			m[val] = true
		} else {
			m[val] = false
		}
	}

	for _, val := range s2 {
		if _, ok := m[val]; ok {
			m[val] = true
		} else {
			m[val] = false
		}
	}

	var cross []int
	for key, value := range m {
		if value {
			cross = append(cross, key)
		}
	}
	return cross
}

func main() {
	cross := crossingSlices2(slice12, slice22)
	fmt.Println("Slice1: ", slice12)
	fmt.Println("Slice2: ", slice22)
	fmt.Println("Crossing: ", cross)
}
