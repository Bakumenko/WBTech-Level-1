package main

import "fmt"

var (
	slice2          = []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	searchedNumber2 = 5
)

func binarySearch2(slice []int, target int) int {
	first := 0
	last := len(slice) - 1
	for first <= last {
		mid := (first + last) / 2
		if target < slice[mid] {
			last = mid - 1
		} else if slice[mid] < target {
			first = mid + 1
		} else if slice[mid] == target {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println("Slice = ", slice2)
	index := binarySearch2(slice2, searchedNumber2)
	fmt.Printf("%v is index of %v", index, searchedNumber2)
}
