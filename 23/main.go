package main

import "fmt"

var (
	slice          = []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	searchedNumber = 5
)

func binarySearch(slice []int, target int, start int, end int) int {
	mid := (start + end) / 2
	if start <= end {
		if target < slice[mid] {
			return binarySearch(slice, target, start, mid-1)
		} else if slice[mid] < target {
			return binarySearch(slice, target, mid+1, end)
		} else if slice[mid] == target {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println("Slice = ", slice)
	index := binarySearch(slice, searchedNumber, 0, len(slice)-1)
	fmt.Printf("%v is index of %v", index, searchedNumber)
}
