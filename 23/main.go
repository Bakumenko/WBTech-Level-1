package main

import "fmt"

var (
	slice1          = []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	searchedNumber1 = 5
)

func binarySearch1(slice []int, target int, start int, end int) int {
	mid := (start + end) / 2
	if start <= end {
		if target < slice[mid] {
			return binarySearch1(slice, target, start, mid-1)
		} else if slice[mid] < target {
			return binarySearch1(slice, target, mid+1, end)
		} else if slice[mid] == target {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println("Slice = ", slice1)
	index := binarySearch1(slice1, searchedNumber1, 0, len(slice1)-1)
	fmt.Printf("%v is index of %v", index, searchedNumber1)
}
